package config

import (
	"io/ioutil"
	"log"
	"sync"
)

type Keys struct {
	Storage *keysStorage
}

type keysStorage struct {
	PrivateKey []byte
	PublicKey  []byte
}

var onceKeys sync.Once
var k *keysStorage

func GetKeysStorage() *keysStorage {
	privateFilePath, publicFlePath := "private.key", "public.key"
	onceKeys.Do(func() {
		pubKey, err := ioutil.ReadFile(publicFlePath)
		if err != nil {
			log.Panicf("Public key generate, error=%v\n", err)
		}

		prvKey, err := ioutil.ReadFile(privateFilePath)
		if err != nil {
			log.Panicf("Private key generate, error=%v\n", err)
		}

		k = &keysStorage{
			PrivateKey: prvKey,
			PublicKey:  pubKey,
		}
	})

	return k
}
