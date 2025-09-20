package main

import (
	"fmt"
	"github.com/bigchain/bigchaind/btcec"
	"github.com/bigchain/bigchaind/btcutil"
	"github.com/bigchain/bigchaind/chaincfg"
)

func main() {
	// Gerar chave privada aleatória
	privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		fmt.Println("Erro ao gerar chave privada:", err)
		return
	}

	// Gerar chave pública
	publicKey := privateKey.PubKey()

	// Gerar endereço usando parâmetros da rede principal
	address, err := btcutil.NewAddressPubKeyHash(
		btcutil.Hash160(publicKey.SerializeCompressed()),
		&chaincfg.MainNetParams,
	)
	if err != nil {
		fmt.Println("Erro ao gerar endereço:", err)
		return
	}

	fmt.Printf("Chave Privada (hex): %x\n", privateKey.Serialize())
	fmt.Printf("Endereço BIGchain: %s\n", address.EncodeAddress())
}
