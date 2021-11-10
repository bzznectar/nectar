// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openapi

import (
	"github.com/gin-gonic/gin"
	"nctr/pkg/service"
	"net/http"
)

type HttpServer struct{}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

type HandlerFunc func(c *gin.Context) error

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(c)
	}
}

func (server *HttpServer) InitHttpServer() {
	gin.SetMode("release")
	r := gin.Default()
	r.Use(Cors())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error":  404,
			"result": nil,
		})
	})
	r.GET("/", indexHandler)
	r.GET("/address", walletAddressHandler)
	r.GET("/chequebook/address", contractAddressHandler)
	r.GET("/exportKey", exportKey)
	r.GET("/harvest/status", mineStatusHandler)
	r.GET("/wallet/balance", walletBalanceHandler)
	r.GET("/chequebook/balance", contractBalanceHandler)

	r.GET("/harvest/deposit", mineDepositHandler)
	r.GET("/harvest/withdraw", mineWithDrawHandler)
	r.GET("/harvest/cashout", cashOutHandler)

	r.GET("/harvest/redeem", undepositHandler)
	r.GET("/status", statusHandler)

	r.Run(service.Node.DebugApi)
}
