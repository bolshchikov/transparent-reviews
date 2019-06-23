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

var PUBLIC = sdk.Export(add, get, getAll)
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
	state.WriteBytes(key, encoded)

	keyString := fmt.Sprintf("%x", hash)

	currentList := state.ReadString(ALL_KEY)
	state.WriteString(ALL_KEY, currentList+","+keyString)

	return keyString
}

func getAll() string {
	return state.ReadString(ALL_KEY)
}

func get(key string) string {
	return string(state.ReadBytes([]byte(key)))
}
