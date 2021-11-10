// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openapi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"nctr/pkg/service"
)

func walletAddressHandler(c *gin.Context) {
	data := WalletData{
		WalletAddress: service.Node.WalletAddress,
	}
	c.JSON(200, gin.H{
		"error":  nil,
		"result": &data,
	})
}
func exportKey(c *gin.Context) {
	privateKey := service.Node.PrivateKey
	privateKeyBytes := crypto.FromECDSA(privateKey)

	priv := common.Bytes2Hex(privateKeyBytes)
	c.JSON(200, gin.H{
		"address":    service.Node.WalletAddress,
		"privateKey": priv,
	})
}

func walletBalanceHandler(c *gin.Context) {

	eth := rpcService.BalanceOfEth()
	nctr := rpcService.BalanceOfNCTR()
	if service.Node.MainNet {
		c.JSON(200, gin.H{
			"xdai": eth,
			"nctr":  nctr,
		})

		return
	}
	c.JSON(200, gin.H{
		"eth": eth,
		"nctr": nctr,
	})

}
