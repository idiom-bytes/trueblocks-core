package unchained

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func simulateClient() (*backends.SimulatedBackend, *bind.TransactOpts, *ecdsa.PrivateKey) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// auth := bind.NewKeyedTransactor(privateKey)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// auth.Nonce = big.NewInt(int64(0))
	// auth.Value = big.NewInt(875000000 * 3) // in wei
	// auth.GasLimit = uint64(300000)         // in units
	// auth.GasPrice = big.NewInt(875000000)

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}
	blockGasLimit := uint64(4712388)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}

	// value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(4712388) // uint64(93556)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return client, auth, privateKey
}

func TestNewUnchained(t *testing.T) {
	client, auth, _ := simulateClient()
	defer client.Close()
	_, tx, instance, err := DeployUnchained(auth, client)
	if err != nil {
		t.Error(err)
	}
	mined := make(chan error)
	go func() {
		_, err = bind.WaitDeployed(context.Background(), client, tx)
		mined <- err
		close(mined)
	}()
	client.Commit()

	select {
	case err = <-mined:
		if err != nil {
			t.Error("transaction error:", err)
		}
	case <-time.After(2 * time.Second):
		t.Errorf("timeout")
	}

	defaultValue, err := instance.ChainNameToHash(nil, "mainnet")
	if err != nil {
		t.Error(err)
	}
	if defaultValue != "QmP4i6ihnVrj8Tx7cTFw4aY6ungpaPYxDJEZ7Vg1RSNSdm" {
		t.Error("Wrong default value", defaultValue)
	}
}

func TestPublishHash(t *testing.T) {
	client, auth, _ := simulateClient()
	// auth.Value = big.NewInt(1000000000000000000) // in wei (1 eth)
	defer client.Close()
	_, tx, instance, err := DeployUnchained(auth, client)
	if err != nil {
		t.Error(err)
	}
	mined := make(chan error)
	go func() {
		_, err = bind.WaitDeployed(context.Background(), client, tx)
		mined <- err
		close(mined)
	}()
	client.Commit()

	select {
	case err = <-mined:
		if err != nil {
			t.Error("transaction error:", err)
		}
	case <-time.After(2 * time.Second):
		t.Errorf("timeout")
	}

	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	publishHashTx, err := instance.PublishHash(auth, "gnosis", "QmP4i6ihnVrj8Tx7cTFw4aY6ungpaPYxDJEZ7Vg1RSNSdm")
	if err != nil {
		t.Error(err)
	}

	mined2 := make(chan error)
	go func() {
		_, err = bind.WaitMined(context.Background(), client, publishHashTx)
		mined2 <- err
		close(mined2)
	}()
	client.Commit()

	select {
	case err = <-mined2:
		if err != nil {
			t.Error("transaction error:", err)
		}
	case <-time.After(2 * time.Second):
		t.Errorf("timeout")
	}

	hashValue, err := instance.ChainNameToHash(nil, "gnosis")
	if err != nil {
		t.Error(err)
	}
	if hashValue != "QmP4i6ihnVrj8Tx7cTFw4aY6ungpaPYxDJEZ7Vg1RSNSdm" {
		t.Error("Wrong default value", hashValue)
	}

}
