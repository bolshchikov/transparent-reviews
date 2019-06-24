package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/env"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(add, getAll)
var SYSTEM = sdk.Export(_init)

var ALL_KEY = []byte("__ALL_REVIEWS_KEY__")

type Review struct {
	Data      string `json:"data"`
	Timestamp uint64 `jspon:"timestamp"`
	Author    []byte `json:"author"`
}

func _init() {}

func add(text string) string {
	hash := md5.Sum([]byte(text))
	key := hash[:]

	if !bytes.Equal(state.ReadBytes(key), nil) {
		panic("Review already exists")
	}

	timestamp := env.GetBlockTimestamp()
	signer := address.GetSignerAddress()
	encoded, _ := json.Marshal(&Review{
		text,
		timestamp,
		signer,
	})
	state.WriteBytes(hash[:], encoded)
	state.WriteBytes(ALL_KEY, append(state.ReadBytes(ALL_KEY), key...))

	return fmt.Sprintf("%x", hash)
}

func getAll() string {
	res := make(map[string]string)
	ids := state.ReadBytes(ALL_KEY)

	idLength := 16

	for i := 0; i < len(ids); i += idLength {
		end := i + idLength

		if end > len(ids) {
			end = len(ids)
		}

		id := ids[i:end]
		res[fmt.Sprintf("%x", id)] = string(state.ReadBytes(id))
	}
	encoding, _ := json.Marshal(&res)
	return string(encoding)
}
