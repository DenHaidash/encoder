package rsa

import (
	mathHelper "encoder/helpers/math"
	b64 "encoding/base64"
	"math/rand"
	"strconv"
	"strings"
)

const keySeparator = ":"

func generateKeys() (*KeyPair, error) {
	const randUpperBound int32 = 100

	q := rand.Int31n(randUpperBound)
	p := rand.Int31n(randUpperBound)

	qs, err := mathHelper.FindClosestPrimeNumber(q)

	if err != nil {
		return nil, err
	}

	ps, err := mathHelper.FindClosestPrimeNumber(p)

	if err != nil {
		return nil, err
	}

	n := ps * qs

	var d, ds int32
	for ds != 1 {
		d = rand.Int31n(randUpperBound)
		ds = mathHelper.Gcd(d, (ps-1)*(qs-1))
	}

	var e, es int32
	for es != 1 {
		e++
		es = (e * d) % ((ps - 1) * (qs - 1))
	}

	keyPair := NewKeyPair(e, d, n)

	return &keyPair, nil
}

func splitKey(key string) (int32, int32, error) {
	decodedKey, err := b64.StdEncoding.DecodeString(key)

	if err != nil {
		return 0, 0, err
	}

	keyParts := strings.Split(string(decodedKey), keySeparator)

	keyPart1, err := strconv.Atoi(keyParts[0])

	if err != nil {
		return 0, 0, err
	}

	keyPart2, err := strconv.Atoi(keyParts[1])

	if err != nil {
		return 0, 0, err
	}

	return int32(keyPart1), int32(keyPart2), nil
}
