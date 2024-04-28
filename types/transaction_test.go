package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlayco/blockverse/crypto"
	"github.com/vlayco/blockverse/proto"
	"github.com/vlayco/blockverse/util"
)

// My balance 100 coins.
// Want to sent 5 coins to "SomeAddress".
// Two outputs are created in order to do this.
// 1. 5 coins to someone we want to send,
// 2. 95 coins back to our address.
// These are the mechanics - entire balance first goes to that address,
// and then the remaining amount is returned back to us.
func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivKey.Public().Address().Bytes()

	toPrivKey := crypto.GeneratePrivateKey()
	toAddress := toPrivKey.Public().Address().Bytes()

	// It tells us about ourselves, what was our previous output we received.
	// It's about ourselves and our previoous receivings.
	// Because this is what we will spend.
	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}

	// Output specifies a destination, where we are going to spend the coins at.
	// This is sending part, with the amount we want to spend.
	// In this hypotetic scenarion lets say that we have 100 of some coins, and
	// we want to send 5 (to some address). But the we also need to spend rest\
	// of the coins we have, 95.
	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}
	output2 := &proto.TxOutput{
		Amount: 95,
		// we are sending back 95 coins to ourselves! In first step all goes,
		// but the change is sent back to us, because we want to spend only
		// 5 coins.
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bytes()

	assert.True(t, VerifyTransaction(tx))

	fmt.Printf("%+v\n", tx)
}
