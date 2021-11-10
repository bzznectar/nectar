// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package openapi

import (
	"github.com/gin-gonic/gin"
	"nctr/pkg/constant"
	"nctr/pkg/store/localstore"
	"math/big"
)

func contractAddressHandler(c *gin.Context) {
	contract := localstore.Get(constant.ContractAddressKey)
	data := ContractData{
		ContractAddress: contract,
	}
	c.JSON(200, gin.H{
		"error":  nil,
		"result": &data,
	})
}

func contractBalanceHandler(c *gin.Context) {
	contractAddress := localstore.Get(constant.ContractAddressKey)

	amount, err := rpcService.BalanceOfMiner(contractAddress)

	if err != nil {
		c.JSON(500, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	incomeData, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return

	}

	temp := big.NewInt(0)
	withdraw := temp.Sub(amount, incomeData.DepositAmount)
	if withdraw.Cmp(big.NewInt(0)) <= 0 {
		withdraw = big.NewInt(0)
	}
	c.JSON(200, gin.H{
		"totalAmount":        amount.String(),
		"pledgeAmount":       incomeData.DepositAmount,
		"withDrawableAmount": withdraw,
	})
}
