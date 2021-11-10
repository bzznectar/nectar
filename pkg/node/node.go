package node

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron/v3"
	"nctr/pkg/abi"
	"nctr/pkg/alog"
	"nctr/pkg/bee"
	"nctr/pkg/constant"
	"nctr/pkg/openapi"
	"nctr/pkg/service"
	"nctr/pkg/store/filestore"
	"nctr/pkg/store/localstore"
	"nctr/pkg/task"
	"nctr/pkg/tools"

	"math/big"
	"os"
)

var (
	c *cron.Cron

	client     *ethclient.Client
	rpcService service.RpcService
)

type Node struct {
}

func (node *Node) InitNode(DataDir string, RpcAddress string, PrivateKey *ecdsa.PrivateKey, PublicKey []byte,
	PssPublicKey []byte,
	OverlayAddress common.Address,
	NetWorkId int,
	WalletAddress string,
	FactoryAddress string,
	SwarmPort string,
	DebugApi string, MainNet bool) {
	service.Node.DataDir = DataDir
	service.Node.RpcAddress = RpcAddress
	service.Node.PrivateKey = PrivateKey
	service.Node.PublicKey = PublicKey
	service.Node.PssPublicKey = PssPublicKey
	service.Node.OverlayAddress = OverlayAddress
	service.Node.NetWorkId = NetWorkId
	service.Node.WalletAddress = WalletAddress
	service.Node.FactoryAddress = FactoryAddress
	service.Node.SwarmPort = SwarmPort
	service.Node.DebugApi = DebugApi
	service.Node.MainNet = MainNet
	localstore.InitDb(DataDir)
	filestore.InitDb(DataDir)
	node.CheckNode()

	var server = openapi.HttpServer{}
	server.InitHttpServer()

}

func (node *Node) CheckNode() {

	DeviceInfo := tools.GetHardwareData()
	if !DeviceInfo.MinimalRequire {
		os.Exit(0)
	}
	service.Node.TotalDiskFree = tools.GetHwDiskFree(DeviceInfo.HW)

	client = rpcService.ConnectRPC()
	amount := rpcService.BalanceOfEth()
	if amount == nil || amount.Cmp(big.NewInt(0)) == 0 {
		alog.Error(" cannot continue until there is sufficient ETH (for Gas)")
		os.Exit(0)
	}
	rpcService.ConnectRPC()
	node.Deploy()
}

func (node *Node) Deploy() {

	contractAddress := localstore.Get(constant.ContractAddressKey)
	contractHash := localstore.Get(constant.ContractHashKey)
	b := bee.Bee{}
	wallet, contract := b.GetBeeNodeInfo("http://localhost" + service.Node.SwarmPort)
	if wallet == "" || contract == "" {
		os.Exit(0)
	}

	if len(contractAddress) > 0 {
		bindContract := localstore.Get(constant.BindSwarmContractKey)
		if contract != bindContract {
			alog.Error("The bee node contract changes")
			os.Exit(0)
		}
		alog.Info("using the chequebook address: " + contractAddress)
		task.InitTask()
		return
	}
	alog.Info("no chequebook found, deploying new one.")
	amount := rpcService.BalanceOfNCTR()
	if amount == nil || amount.Cmp(big.NewInt(constant.PledgeAmount)) < 0 {
		alog.Error(" cannot continue until there is  at least ", constant.PledgeAmount/10000000000000000, " NCTR")

		os.Exit(0)
	}

	if len(contractHash) > 0 {
		bindContract := localstore.Get(constant.BindSwarmContractKey)
		if contract != bindContract {
			alog.Error("The bee node contract changes")
			os.Exit(0)
		}
		alog.Info(" using the chequebook transaction hash: " + contractHash)
		node.checkBlockTask()
		return
	}
	if rpcService.SwarmContractBindStatus(contract) {
		alog.Error("The Bee node has been bound")
		return
	}

	token, err := abi.NewFactory(common.HexToAddress(constant.FactoryAddress), client)

	if err != nil {
		alog.Error(err.Error())
		os.Exit(0)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		alog.Error(err.Error())
		os.Exit(0)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(service.Node.PrivateKey, chainId)
	if err != nil {
		alog.Error(err.Error())
		os.Exit(0)
	}
	alog.Info(service.Node.WalletAddress)
	alog.Info(contract)
	t, err := token.DeployMiner(opts, common.HexToAddress(service.Node.WalletAddress), common.HexToAddress(contract))
	if err != nil {
		alog.Error(err.Error())
		os.Exit(0)
	}
	hash := t.Hash()
	alog.Info("successfully deposited to chequebook")
	localstore.Put(constant.ContractHashKey, hash.String())
	localstore.Put(constant.BindSwarmContractKey, contract)
	alog.Info("using the chequebook transaction hash " + hash.String())
	node.checkBlockTask()

}

func (node *Node) checkBlockTask() {

	c = cron.New(cron.WithSeconds())
	c.AddFunc("*/20 * * * * ?", node.checkBlockStatus)
	c.Start()

}

func (node *Node) checkBlockStatus() {
	txHash := localstore.Get(constant.ContractHashKey)
	query := ethereum.FilterQuery{FromBlock: nil, ToBlock: nil, Addresses: []common.Address{common.HexToAddress(constant.FactoryAddress)}}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		alog.Error(err.Error())
		os.Exit(0)
	}
	flag := false
	address := ""
	for _, log := range logs {
		if txHash == log.TxHash.String() {
			addr := common.Bytes2Hex(log.Data)
			addr = "0x" + string(addr[len(addr)-40:])

			localstore.Put(constant.ContractAddressKey, addr)

			alog.Info("using the chequebook address: " + addr)
			flag = true
			address = addr
			c.Stop()
			break
		}

	}

	if flag {
		node.recharge(address)
	}

}

func (node *Node) recharge(addr string) {
	amount := rpcService.BalanceOfNCTR()
	t, err := rpcService.Recharge(amount, addr)
	if err != nil {
		alog.Error(err.Error())
		return
	}
	alog.Info("sent deposit transaction " + t.Hash().String())
	bind.WaitMined(context.Background(), client, t)
	alog.Info("successfully deposited to chequebook")
	task.InitTask()
}
