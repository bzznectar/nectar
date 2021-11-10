// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openapi

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"nctr/pkg/constant"
	"nctr/pkg/service"
	"nctr/pkg/store/localstore"
	"nctr/pkg/version"
	"math/big"
)

// @Summary Welcom
// @Accept  json
// @Produce json
// @Success 200 string  "Welcome to nctr"
// @Router / [get]
func indexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Welcome to nctr",
	})
}

func statusHandler(c *gin.Context) {
	contractAddress := localstore.Get(constant.ContractAddressKey)
	walletAddress := service.Node.WalletAddress

	info, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {

		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	amount, err := rpcService.BalanceOfMiner(contractAddress)
	if err != nil {

		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	temp := big.NewInt(0)
	withdraw := temp.Sub(amount, info.DepositAmount)
	if withdraw.Cmp(big.NewInt(0)) <= 0 {
		withdraw = big.NewInt(0)
	}
	data := StatusData{
		Status: 1, WalletAddress: walletAddress,
		Version:         version.Version,
		ContractAddress: contractAddress,
		PledgeBalance:   info.DepositAmount,
		TotalBalance:    amount,
		RewardBalance:   withdraw,
		OverlayAddress:  service.Node.OverlayAddress.String(),
		PublicKey:       hexutil.Encode(service.Node.PublicKey),
		PssPublicKey:    hexutil.Encode(service.Node.PssPublicKey),
	}
	c.JSON(200, gin.H{
		"result": data,
		"error":  nil,
	})

}
