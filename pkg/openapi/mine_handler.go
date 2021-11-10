// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openapi

import (
	"github.com/gin-gonic/gin"
	"math/big"
	"nctr/pkg/constant"
	"nctr/pkg/service"
	"nctr/pkg/store/localstore"
	"os"
	"strconv"
	"time"
)

var rpcService service.RpcService

func mineStatusHandler(c *gin.Context) {
	contractAddress := localstore.Get(constant.ContractAddressKey)
	res, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"harvest": service.Node.MinerStatus,
		"reward":  res.AllReturnAmount,
		"pending": res.RewardAmount,
		"pledge":  res.DepositAmount,
	})
}

func mineDepositHandler(c *gin.Context) {

	amount := c.DefaultQuery("amount", strconv.FormatInt(constant.PledgeAmount, 10))
	n, _ := strconv.ParseInt(amount, 10, 64)

	if n < constant.PledgeAmount {
		c.JSON(200, gin.H{
			"error":  "The number of NCTR needs to be greater than " + strconv.FormatInt(constant.PledgeAmount, 10),
			"result": nil,
		})
		return
	}

	contractAddress := localstore.Get(constant.ContractAddressKey)
	incomeData, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	if incomeData.DepositAmount.Cmp(big.NewInt(0)) > 0 {
		c.JSON(200, gin.H{
			"error":  "You have pledged",
			"result": nil,
		})
		return
	}

	num, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	balance := rpcService.BalanceOfNCTR()

	if balance.Cmp(big.NewInt(num)) < 0 {
		c.JSON(200, gin.H{
			"error":  "nctr insufficient balance",
			"result": nil,
		})
		return
	}

	t, err := rpcService.Recharge(big.NewInt(num), contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"error":  nil,
		"result": t.Hash().String(),
	})

}

func mineWithDrawHandler(c *gin.Context) {

	contractAddress := localstore.Get(constant.ContractAddressKey)
	max, err := rpcService.BalanceOfMiner(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	incomeData, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {
		c.JSON(500, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
	}
	pledgeNum := incomeData.DepositAmount
	temp := big.Int{}
	num := temp.Sub(max, pledgeNum)
	if num.Cmp(big.NewInt(0)) <= 0 {
		c.JSON(200, gin.H{
			"error":  "Insufficient cash balance",
			"result": nil,
		})
		return
	}

	t, err := rpcService.WithDraw(contractAddress, num)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"error":  nil,
		"result": t.Hash().String(),
	})

}

func cashOutHandler(c *gin.Context) {
	contractAddress := localstore.Get(constant.ContractAddressKey)
	amount := rpcService.GetRewardAmount(contractAddress)
	if amount == nil || amount.Cmp(big.NewInt(0)) == 0 {
		c.JSON(200, gin.H{
			"error":  "No dividends",
			"result": nil,
		})
		return
	}

	gasAmount := rpcService.GetGasAmount()
	t, err := rpcService.GetReward(amount, contractAddress, gasAmount)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"error":  nil,
		"result": t.Hash().String(),
	})

}

func undepositHandler(c *gin.Context) {
	contractAddress := localstore.Get(constant.ContractAddressKey)
	incomeData, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	if incomeData.DepositAmount.Cmp(big.NewInt(0)) == 0 {
		c.JSON(400, gin.H{
			"error":  "Please pledge first",
			"result": nil,
		})
		return
	}
	t, err := rpcService.UnDeposit(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"error":  nil,
		"result": t.Hash().String(),
	})
	go func() {
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
}
