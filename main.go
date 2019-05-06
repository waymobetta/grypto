package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/waymobetta/wmb"
)

var (
	INFURA_API_KEY = os.Getenv("PERSONAL_INFURA_API_KEY")
	// Infura Rinkeby URL
	INFURA_RINKEBY = fmt.Sprintf("https://rinkeby.infura.io/v3/%s", INFURA_API_KEY)
	// Infura Mainnet URL
	INFURA_MAINNET = fmt.Sprintf("https://mainnet.infura.io/v3/%s", INFURA_API_KEY)
	// Local Node
	LOCAL_NODE = "http://localhost:8545"
	// Ethereum Client endpoint
	ETHEREUM_CLIENT_URL = LOCAL_NODE
)

func main() {
	wmb.Clear()

	genHeader()

	if len(os.Args) < 2 {
		fmt.Println("usage: grypto <number_keys>\n")
		os.Exit(1)
	}

	numKeys, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("[usi] error determing number of keys to produce: %v\n", err)
	}

	t := time.Now()

	// log.Println("[ethereum] connecting..")

	client, err := ethclient.Dial(ETHEREUM_CLIENT_URL)
	if err != nil {
		log.Fatalf("[ethereum] error connecting to client: %v\n", err)
	}

	log.Println("[grypto] initializing key generation sequence..")

	var accountsFound int

	for i := 0; i < numKeys; i++ {
		acct := new(Account)

		err := acct.GenRandKey()
		if err != nil {
			log.Fatalf("[crypto] error generating key: %v\n", err)
		}

		err = acct.Unlock()
		if err != nil {
			log.Fatalf("[unlock] error unlocking account: %v\n", err)
		}

		err = acct.GetBalance(client)
		if err != nil {
			log.Fatalf("[balance] error getting balance: %v\n", err)
		}

		if acct.Balance.Int64() > int64(0) {
			err := acct.Log()
			if err != nil {
				log.Fatalf("[log] error logging account: %v\n", err)
			}
			accountsFound += 1
		}
		// time.Sleep(time.Millisecond * 250)
	}
	log.Printf("[time] took %v seconds", time.Since(t).Seconds())
	log.Printf("[grypto] keys found: %v\n", accountsFound)
}
