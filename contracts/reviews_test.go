package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
)

func TestAddingRecrod(t *testing.T) {
	InServiceScope(nil, nil, func(m Mockery) {
		text := "hello world"
		hash := md5.Sum([]byte(text))
		require.EqualValues(t, add(text), fmt.Sprintf("%x", hash))
	})
}

func TestAddAndGetRecord(t *testing.T) {
	InServiceScope(nil, nil, func(m Mockery) {
		text := "hello world"
		id := add(text)
		messages := getAll()
		var res map[string]string
		json.Unmarshal([]byte(messages), &res)
		require.Contains(t, res[id], text)
	})
}
