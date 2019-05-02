package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	Address common.Address `json:"address"`
	Private string         `json:"private"`
	Balance *big.Int       `json:"balance"`
}
