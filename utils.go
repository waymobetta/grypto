package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (a *Account) Log() error {
	accountFile := "accounts.json"

	accountSlice := []*Account{}

	accountBytes, err := ioutil.ReadFile(accountFile)
	if err != nil {
		return err
	}

	if len(accountBytes) > 0 {
		err = json.Unmarshal(accountBytes, &accountSlice)
		if err != nil {
			return err
		}
	}

	accountSlice = append(accountSlice, a)

	accountBytes, err = json.MarshalIndent(&accountSlice, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(accountFile, accountBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) Unlock() error {
	privateKey, err := crypto.HexToECDSA(a.Private)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	a.Address = fromAddress

	return nil
}

func (a *Account) GetBalance(client *ethclient.Client) error {
	ctx := context.Background()

	weiBalance, err := client.BalanceAt(ctx, a.Address, nil)
	if err != nil {
		return err
	}

	ethBalance := weiBalance.Div(weiBalance, big.NewInt(1000000000000000000))

	a.Balance = ethBalance

	return nil
}

func genRandKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func genHeader() {
	fmt.Println("grypto\n")
}
