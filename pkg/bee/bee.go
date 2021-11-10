// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bee

import (
	"bytes"
	"encoding/json"
	"nctr/pkg/alog"
	"io"
	"net/http"
	"os"
	"time"
)

type Bee struct {
	WalletAddress   string
	ContractAddress string
}

type WalletAddressData struct {
	Overlay      string `json:"overlay"`
	Ethereum     string `json:"ethereum"`
	PssPublicKey string `json:"pssPublicKey"`
	PublicKey    string `json:"publicKey"`
}

type ContractAddressData struct {
	ChequebookAddress string `json:"chequebookAddress"`
}

func (bee *Bee) GetWalletAddress(url string, c chan string) {
	res, err := bee.GetRequest(url + "/addresses")
	if err != nil {
		alog.Error(err.Error())
		c <- ""
		return
	}
	var wallet WalletAddressData
	err = json.Unmarshal(res, &wallet)
	if err != nil {
		alog.Error(err.Error())
		os.Exit(0)
		c <- ""
		return
	}

	c <- wallet.Ethereum

}

func (bee *Bee) GetRequest(url string) ([]byte, error) {
	client := http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

	}

	return result.Bytes(), nil

}

func (bee *Bee) GetContractAddress(url string, c chan string) {
	res, err := bee.GetRequest(url + "/chequebook/address")
	if err != nil {
		alog.Error(err.Error())
		c <- ""
		return
	}
	var contract ContractAddressData
	err = json.Unmarshal(res, &contract)
	if err != nil {
		alog.Error(err.Error())
		c <- ""
		return
	}
	c <- contract.ChequebookAddress

}

func (bee *Bee) GetBeeNodeInfo(url string) (string, string) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	go bee.GetWalletAddress(url, ch1)
	go bee.GetContractAddress(url, ch2)
	wallet, contract := <-ch1, <-ch2

	return wallet, contract
}
