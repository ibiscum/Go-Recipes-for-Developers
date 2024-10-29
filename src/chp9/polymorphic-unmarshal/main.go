package main

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
)

type KeyType string

const (
	KeyTypeRSA     = "rsa"
	KeyTypeED25519 = "ed25519"
)

type Key struct {
	Type KeyType          `json:"type"`
	Key  crypto.PublicKey `json:"key"`
}

var (
	KeyUnmarshalers = map[KeyType]func(json.RawMessage) (crypto.PublicKey, error){}
)

func RegisterKeyUnmarshaler(keyType KeyType, unmarshaler func(json.RawMessage) (crypto.PublicKey, error)) {
	KeyUnmarshalers[keyType] = unmarshaler
}

var ErrInvalidKeyType = errors.New("Invalid key type")

func (k *Key) UnmarshalJSON(in []byte) error {
	type keyUnmarshal struct {
		Type KeyType         `json:"type"`
		Key  json.RawMessage `json:"key"`
	}

	var key keyUnmarshal
	err := json.Unmarshal(in, &key)
	if err != nil {
		return err
	}
	k.Type = key.Type
	unmarshaler := KeyUnmarshalers[key.Type]
	if unmarshaler == nil {
		return ErrInvalidKeyType
	}
	k.Key, err = unmarshaler(key.Key)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	RegisterKeyUnmarshaler(KeyTypeRSA, func(in json.RawMessage) (crypto.PublicKey, error) {
		var key rsa.PublicKey
		if err := json.Unmarshal(in, &key); err != nil {
			return nil, err
		}
		return &key, nil
	})
	RegisterKeyUnmarshaler(KeyTypeED25519, func(in json.RawMessage) (crypto.PublicKey, error) {
		var key ed25519.PublicKey
		if err := json.Unmarshal(in, &key); err != nil {
			return nil, err
		}
		return &key, nil
	})

	input := []byte(`[
{
  "type": "rsa",
  "key": {"N": 123,"E":456}
},
{
  "type": "ed25519",
  "key": [0,0,0,0,0]
}
]`)
	var keys []Key
	if err := json.Unmarshal(input, &keys); err != nil {
		panic(err)
	}

	fmt.Println(keys)
}
