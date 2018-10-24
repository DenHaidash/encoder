package math

import "errors"

func SundaramPrimes(n int32) []int32 {
	primesMatrix := make([]int8, n)

	for i := int32(1); n > 3*i+1; i++ {
		for j := int32(1); n > i+j+2*i*j && j <= i; j++ {
			primesMatrix[i+j+2*i*j] = 1
		}
	}

	primes := make([]int32, 0)

	for i, v := range primesMatrix {
		if v == 0 {
			primes = append(primes, int32(2*i+1))
		}
	}

	return primes
}

func FindClosestPrimeNumber(n int32) (int32, error) {
	primes := SundaramPrimes(n)

	if len(primes) == 0 {
		return 0, errors.New("Primes not found")
	}

	return primes[len(primes)-1], nil
}
