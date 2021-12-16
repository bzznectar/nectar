package task

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron/v3"
	"math/big"
	"nctr/pkg/abi"
	"nctr/pkg/alog"
	"nctr/pkg/constant"
	"nctr/pkg/service"
	"nctr/pkg/store/filestore"
	"nctr/pkg/store/localstore"
	"nctr/pkg/tools"
	"os"
	"strconv"
	"time"
)

var (
	client   *ethclient.Client
	blockNum uint64 = 0
	IdKey           = "id_key"
	MaxId           = 100000
)
var rpcService service.RpcService

type HeartData struct {
	Status    int
	Timestamp int64
	Info      string
}

func InitTask() {
	levelValue := localstore.Get(constant.WriteLevelKey)
	if levelValue == "" && tools.DeviceInfo.HW > 1 {

		hw := tools.DeviceInfo.HW
		fmt.Println("Which hardware options for harvest:")
		switch hw {

		case 2:
			fmt.Println("1. CPU(4core) RAM(4G) Storage(50G)")
			fmt.Println("2. CPU(8core) RAM(8G) Storage(450G)")
			service.Node.TotalDiskFree = tools.MediumHW.FreeDisk
			service.Node.Hw = 3
			break

		case 3:
			fmt.Println("1. CPU(4core) RAM(4G) Storage(50G)")
			fmt.Println("2. CPU(8core) RAM(8G) Storage(450G)")
			fmt.Println("3. CPU(16core) RAM(16G) Storage(950G)")
			service.Node.TotalDiskFree = tools.HighHW.FreeDisk
			service.Node.Hw = 2
			break

		case 1:
		default:
			fmt.Println("1. CPU(4core) RAM(4G) Storage(50G)")
			service.Node.TotalDiskFree = tools.LowHW.FreeDisk
			service.Node.Hw = 1
			break

		}
		var level string
		fmt.Scan(&level)
		localstore.Put(constant.WriteLevelKey, level)
	} else {
		switch levelValue {
		case "1":

			service.Node.TotalDiskFree = tools.LowHW.FreeDisk
			service.Node.Hw = 1
			break
		case "2":

			service.Node.TotalDiskFree = tools.MediumHW.FreeDisk
			service.Node.Hw = 2
			break
		case "3":

			service.Node.TotalDiskFree = tools.HighHW.FreeDisk
			service.Node.Hw = 3
			break
		default:
			service.Node.TotalDiskFree=tools.LowHW.FreeDisk
			service.Node.Hw = 1


		}

	}

	client = rpcService.ConnectRPC()
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */5 * * * ?", heart)
	//c.AddFunc("*/5 * * * * ?", heart)
	go checkBlock()
	go createFileJob()
	go c.Start()

}

func checkBlock() {
	for {

		num, _ := client.BlockNumber(context.Background())
		if num > blockNum {
			for {
				blockNum += 1
				if blockNum == num {
					break
				}
			}
			if service.Node.MinerStatus {
				alog.Info("harvest service: updated block height to " + strconv.FormatUint(num, 10))
			} else {
				alog.Info("File server : p2p files is syncing,updated block height to " + strconv.FormatUint(num, 10))
			}

		}
		v := tools.GetRandomInt(5, 10)
		time.Sleep(time.Duration(v) * time.Second)

	}

}
func createFileJob() {
	size := float64(service.Node.TotalDiskFree) * constant.WriteDiskCoefficient

	MaxId = int(size / (1024 * 1024 * 4))

	_id := 1
	id := filestore.Get(IdKey)
	if id != "" {
		_id, _ = strconv.Atoi(id)
		_id += 1

	}
	for i := _id; i < MaxId; i++ {
		createFiles(i)
	}

	service.Node.MinerStatus = true

}
func createFiles(id int) {
	//str := tools.GetRandomString(int(tools.GetRandomInt(3*1024*1024, 4*1024*1024)))
	str := tools.GetRandomString(4 * 1024 * 1024)
	filestore.Put(strconv.Itoa(id), str)
	filestore.Put(IdKey, strconv.Itoa(id))

}

func heart() {

	tools.SetBeePort(service.Node.SwarmAddr)

	DeviceInfo := tools.GetHardwareData()
	service.Node.TotalDiskFree = tools.GetHwDiskFree(DeviceInfo.HW)
	if !DeviceInfo.MinimalRequire {
		os.Exit(0)
	}
	if !service.Node.MinerStatus {
		return
	}
	amount := rpcService.BalanceOfEth()
	if amount == nil || amount.Cmp(big.NewInt(0)) == 0 {
		alog.Error("Insufficient eth balance")
		os.Exit(0)
	}
	contractAddress := localstore.Get(constant.ContractAddressKey)

	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		alog.Error(err.Error())
		os.Exit(0)
	}
	opts := bind.CallOpts{From: common.HexToAddress(service.Node.WalletAddress)}
	amount, err = token.GetBalance(&opts)
	if err != nil {
		alog.Error(err.Error())
		os.Exit(0)
	}
	if amount == nil || amount.Cmp(big.NewInt(0)) == 0 {
		alog.Error("Please recharge nctr to the contract first")
		return
	}

	alog.Info("harvest service: harvest packing data ...")
	timestamp := time.Now().Unix()
	key := constant.HeartKeyPrefix + strconv.FormatInt(timestamp, 10)
	b, err := json.Marshal(DeviceInfo)
	if err != nil {
		alog.Error(err.Error())
		return
	}
	heartData := HeartData{Timestamp: timestamp, Info: string(b), Status: 0}
	valid := DeviceInfo.Valid
	if valid {
		heartData.Status = 1
	}
	h, err := json.Marshal(heartData)
	if err != nil {
		alog.Error(err.Error())
		return
	}

	localstore.Put(key, string(h))
	last := localstore.Get(constant.LastUploadTimeKey)
	result := localstore.FindAllHeart(last)
	//start
	if len(result) >= constant.UploadNum {

		chainId, err := client.ChainID(context.Background())
		if err != nil {
			alog.Error(err.Error())
			return
		}
		opts, err := bind.NewKeyedTransactorWithChainID(service.Node.PrivateKey, chainId)
		opts.From = common.HexToAddress(service.Node.WalletAddress)
		price, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			alog.Error(err.Error())
			return
		}
		msg := ethereum.CallMsg{From: common.HexToAddress(service.Node.WalletAddress), GasPrice: price}
		limit, err := client.EstimateGas(context.Background(), msg)
		if err != nil {
			alog.Error(err.Error())
			return
		}
		opts.GasLimit = limit * 100

		r, err := json.Marshal(heartData)
		if err != nil {
			alog.Error(err.Error())
			return
		}
		result := tools.PswEncrypt(string(r))
		sign := getSign(timestamp)

		t, err := token.UploadInfo(opts, big.NewInt(timestamp), strconv.FormatInt(timestamp, 10), sign, result, big.NewInt(int64(service.Node.Hw)))
		if err != nil {
			alog.Error(err.Error())
			return
		}
		alog.Info("harvest service: harvest " + contractAddress + " harvest: true  height: " + strconv.FormatInt(int64(blockNum), 10) + ",tx " + t.Hash().String())

		localstore.Put(constant.HeartHashKeyPrefix, t.Hash().String())
		localstore.Put(constant.LastUploadTimeKey, strconv.FormatInt(timestamp, 10))
	}
	//end

}

func getSign(timestamp int64) [32]byte {
	return sha256.Sum256([]byte(constant.EncryptKey + strconv.FormatInt(timestamp, 10)))
}

