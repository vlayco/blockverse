package types

import (
	"crypto/sha256"

	"github.com/vlayco/blockverse/crypto"
	"github.com/vlayco/blockverse/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) *crypto.Signature {
	return pk.Sign(HashTransaction(tx))
}

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}

func VerifyTransaction(tx *proto.Transaction) bool {
	for _, input := range tx.Inputs {
		var (
			sig    = crypto.SignatureFromBytes(input.Signature)
			pubKey = crypto.PublicKeyFromBytes(input.PublicKey)
		)
		// This is because before signing transaction we don't have the signature!
		// TODO: make shure we don't run into problems after verification because
		// we've set the signature to nil!
		input.Signature = nil
		if !sig.Verify(pubKey, HashTransaction(tx)) {
			return false
		}
	}
	return true
}
