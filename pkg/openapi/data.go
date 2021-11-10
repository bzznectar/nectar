// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openapi

import "math/big"

type WalletData struct {
	WalletAddress string `json:"address"`
}
type ContractData struct {
	ContractAddress string `json:"chequebook"`
}
type StatusData struct {
	Version         string   `json:"version"`
	Status          int      `json:"status"`
	PublicKey       string   `json:"publicKey"`
	OverlayAddress  string   `json:"overlayAddress"`
	PssPublicKey    string   `json:"pssPublicKey"`
	WalletAddress   string   `json:"walletAddress"`
	ContractAddress string   `json:"contractAddress"`
	TotalBalance    *big.Int `json:"totalBalance"`
	PledgeBalance   *big.Int `json:"pledgeBalance"`
	RewardBalance   *big.Int `json:"rewardBalance"`
}
