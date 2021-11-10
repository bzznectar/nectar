// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"nctr/pkg/abi"
	"nctr/pkg/alog"
	"nctr/pkg/constant"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type RpcService struct {
}
type NodeData struct {
	DataDir        string
	RpcAddress     string
	PrivateKey     *ecdsa.PrivateKey
	PublicKey      []byte
	PssPublicKey   []byte
	OverlayAddress common.Address
	NetWorkId      int
	WalletAddress  string
	FactoryAddress string
	SwarmPort      string
	DebugApi       string
	MainNet        bool
	MinerStatus    bool
	TotalDiskFree  uint64
}

var (
	Node   NodeData
	client *ethclient.Client
)

func (service *RpcService) ConnectRPC() *ethclient.Client {
	if client == nil {
		rpcClient, err := rpc.Dial(Node.RpcAddress)
		if err != nil {
			alog.Error(err.Error())
			os.Exit(1)
		}
		conn := ethclient.NewClient(rpcClient)
		client = conn
	}

	return client

}

type WalletData struct {
	Address string `json:"address"`
	Crypto  string `json:"crypto"`
	Salt    string `json:"salt"`
	Version string `json:"version"`
}
type IncomeData struct {
	DepositAmount   *big.Int
	AllRewardAmount *big.Int
	AllReturnAmount *big.Int
	RewardAmount    *big.Int
}

func (service *RpcService) BalanceOfEth() *big.Int {

	balance, err := client.BalanceAt(context.TODO(), common.HexToAddress(Node.WalletAddress), nil)
	if err != nil {
		alog.Error(err)
		return nil
	}
	return balance

}

func (service *RpcService) BalanceOfNCTR() *big.Int {
	client = service.ConnectRPC()
	token, err := abi.NewNctr(common.HexToAddress(constant.HctrAddress), client)
	if err != nil {
		return nil
	}
	amount, err := token.BalanceOf(nil, common.HexToAddress(Node.WalletAddress))
	if err != nil {
		return nil
	}
	return amount
}

func (service *RpcService) BalanceOfMiner(contractAddress string) (*big.Int, error) {
	client = service.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress)}
	amount, err := token.GetBalance(&opts)
	if err != nil {
		return nil, err
	}

	return amount, nil
}

func (service *RpcService) HinDecimal() uint8 {
	client = service.ConnectRPC()
	token, err := abi.NewNctr(common.HexToAddress(constant.HctrAddress), client)
	if err != nil {
		return 0
	}
	amount, err := token.Decimals(nil)
	if err != nil {
		return 0
	}

	return amount
}

func (service *RpcService) GetRewardAmount(address string) *big.Int {
	client = service.ConnectRPC()
	token, err := abi.NewPool(common.HexToAddress(constant.PoolAddress), client)
	if err != nil {
		alog.Error(err.Error())
		return nil
	}
	amount, err := token.GetRewardAmount(nil, common.HexToAddress(address))
	if err != nil {
		alog.Error(err.Error())
		return nil
	}
	return amount

}

func (service *RpcService) GetGasAmount() *big.Int {
	client = service.ConnectRPC()
	token, err := abi.NewPool(common.HexToAddress(constant.PoolAddress), client)
	amount, err := token.FeeAmount(nil)
	if err != nil {
		alog.Error(err.Error())
		return nil
	}
	return amount

}

func (service *RpcService) GetReward(amount *big.Int, contractAddress string, gas *big.Int) (*types.Transaction, error) {
	client = service.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {

		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	opts.Value = gas
	opts.From = common.HexToAddress(Node.WalletAddress)

	t, err := token.GetReward(opts, amount)
	if err != nil {
		return nil, err

	}

	return t, nil

}

func (service *RpcService) WithDraw(contractAddress string, amount *big.Int) (*types.Transaction, error) {
	client = service.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {

		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {

		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {

		return nil, err
	}
	t, err := token.Withdraw(opts, amount)
	if err != nil {

		return nil, err
	}
	return t, nil

}

func (service *RpcService) GetAllInfo(address string) (*IncomeData, error) {
	client = service.ConnectRPC()
	token, err := abi.NewPool(common.HexToAddress(constant.PoolAddress), client)
	if err != nil {
		return nil, err
	}
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress)}
	res, err := token.GetAllInfo(&opts, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	var data IncomeData
	data.DepositAmount = res.DepositAmount
	data.AllRewardAmount = res.AllRewardAmount
	data.AllReturnAmount = res.AllReturnAmount
	data.RewardAmount = res.RewardAmount
	return &data, nil
}

func (service *RpcService) UnDeposit(contractAddress string) (*types.Transaction, error) {
	client = service.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)

	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.HexToAddress(Node.WalletAddress), GasPrice: price}
	limit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	opts.GasLimit = limit * 100

	t, err := token.Cashout(opts)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (service *RpcService) Recharge(amount *big.Int, contractAddress string) (*types.Transaction, error) {
	client = service.ConnectRPC()
	nctr, err := abi.NewNctr(common.HexToAddress(constant.HctrAddress), client)
	chainId, err := client.ChainID(context.Background())
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	t, err := nctr.Approve(opts, common.HexToAddress(contractAddress), amount)
	bind.WaitMined(context.Background(), client, t)
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	t, err = token.Deposit(opts, amount)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (service *RpcService) SwarmContractBindStatus(swarmContract string) bool {
	client = service.ConnectRPC()
	token, err := abi.NewPool(common.HexToAddress(constant.PoolAddress), client)
	if err != nil {
		alog.Error(err.Error())
		return true

	}

	addr, err := token.GetNCTRContract(nil, common.HexToAddress(swarmContract))
	if err != nil {
		alog.Error(err.Error())
		return true

	}
	alog.Info("addr: " + addr.String())
	if addr.String() == "0x0000000000000000000000000000000000000000" {
		return false
	}
	return true

}
