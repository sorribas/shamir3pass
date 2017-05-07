package shamir3pass

import (
	"crypto/rand"
	"math/big"
)

type Key struct {
	Encryption *big.Int
	Decryption *big.Int
	Prime      *big.Int
}

func GenerateKey(prime *big.Int) Key {
	if prime == nil {
		prime = Random1024BitPrime()
	}
	for {
		n := random2048()
		primeMinusOne := &big.Int{}
		primeMinusOne.Sub(prime, big.NewInt(1))
		gcd := &big.Int{}
		gcd.GCD(nil, nil, n, primeMinusOne)
		if gcd.Cmp(big.NewInt(1)) == 0 {
			mi := &big.Int{}
			mi.ModInverse(n, primeMinusOne)
			return Key{
				Encryption: n,
				Decryption: mi,
				Prime:      prime,
			}
		}
	}
}

func random2048() *big.Int {
	// (2^2048 - 1) - 2 ^ 2047
	twoToThe47th := &big.Int{}
	twoToThe47th.Exp(big.NewInt(2), big.NewInt(2047), nil)

	size := &big.Int{}
	size.Exp(big.NewInt(2), big.NewInt(2048), nil)
	size.Sub(size, big.NewInt(1))
	size.Sub(size, twoToThe47th)
	random, err := rand.Int(rand.Reader, size)
	if err != nil {
		panic(err)
	}
	random.Add(random, twoToThe47th)
	return random
}

func Random1024BitPrime() *big.Int {
	prime, err := rand.Prime(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	return prime
}

func Encrypt(m *big.Int, key Key) *big.Int {
	r := &big.Int{}
	r.Exp(m, key.Encryption, key.Prime)
	return r
}

func Decrypt(c *big.Int, key Key) *big.Int {
	r := &big.Int{}
	r.Exp(c, key.Decryption, key.Prime)
	return r
}
