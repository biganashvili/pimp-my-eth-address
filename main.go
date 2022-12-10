package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	var prefix string
	var suffix string
	var contains string

	// flags declaration using flag package
	flag.StringVar(&prefix, "p", "", "Specify username. Default is root")
	flag.StringVar(&suffix, "s", "", "Specify pass. Default is password")
	flag.StringVar(&contains, "c", "", "Specify pass. Default is password")

	flag.Parse() // after declaring flags we need to call it

	// check if cli params match

	for {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		// fmt.Println("Public Key:", hexutil.Encode(publicKeyBytes))

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		// 0xB10C0DE
		if prefix != "" && !strings.HasPrefix(address, "0x"+prefix) {
			continue
		}
		if suffix != "" && !strings.HasSuffix(address, suffix) {
			continue
		}
		if contains != "" && !strings.Contains(address, contains) {
			continue
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		fmt.Println("Horray!:")
		fmt.Println("Address:", address)
		fmt.Println("prv_key:", hexutil.Encode(privateKeyBytes))
	}
}
