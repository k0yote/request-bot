package bot

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func reqBody() ([]byte, error) {
	u := map[string]interface{}{}

	uid, err := generateUUID()
	if err != nil {
		return nil, err
	}

	addr, err := generateOnetimeSigner()
	if err != nil {
		return nil, err
	}

	u["uuid"] = hexutil.Encode(uid)
	u["orderId"] = uuid.New()
	u["mintInfo"] = map[string]interface{}{
		"mintData": map[string]interface{}{"walletAddress": addr.Hex()},
	}

	return json.Marshal(u)
}

func generateUUID() ([]byte, error) {
	generatedId := uuid.New()
	id := [32]byte{}
	for i := 0; i < 16; i++ {
		id[16+i] = generatedId[i]
	}

	return hexutil.Decode(hexutil.Encode(id[16:]))
}

func generateOnetimeSigner() (common.Address, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("failed to decode public key from private key, maybe wrong format")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}
