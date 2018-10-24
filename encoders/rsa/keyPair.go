package rsa

import (
	b64 "encoding/base64"
	"strconv"
)

type KeyPair struct {
	E, D, N    int32
	PublicKey  string
	PrivateKey string
}

func NewKeyPair(e, d, n int32) KeyPair {
	return KeyPair{
		E:          e,
		D:          d,
		N:          n,
		PublicKey:  b64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(n)) + keySeparator + strconv.Itoa(int(d)))),
		PrivateKey: b64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(n)) + keySeparator + strconv.Itoa(int(e)))),
	}
}
