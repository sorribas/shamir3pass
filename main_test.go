package shamir3pass

import (
	"fmt"
	"math/big"
	"testing"
)

func TestBasic(t *testing.T) {
	key := GenerateKey(1024)
	message := big.NewInt(1331)

	crypto := Encrypt(message, key)
	decrypted := Decrypt(crypto, key)
	fmt.Printf("%v\n", decrypted)

	if decrypted.String() != "1331" {
		t.Fail()
	}
}

func TestDoubleEncryption(t *testing.T) {
	key := GenerateKey(1024)
	key2 := GenerateKeyFromPrime(key.Prime)

	message := big.NewInt(1331)

	crypto := Encrypt(message, key)
	crypto2 := Encrypt(crypto, key2)
	crypto3 := Decrypt(crypto2, key)
	decrypted := Decrypt(crypto3, key2)

	fmt.Printf("%v\n", decrypted)

	if decrypted.String() != "1331" {
		t.Fail()
	}
}

func Test4keys(t *testing.T) {
	key1 := GenerateKey(1024)
	key2 := GenerateKeyFromPrime(key1.Prime)
	key3 := GenerateKeyFromPrime(key1.Prime)
	key4 := GenerateKeyFromPrime(key1.Prime)

	fmt.Printf("BIG: %v\n\n", key1.Encryption)
	message := big.NewInt(1331)

	crypto1 := Encrypt(message, key1)
	crypto2 := Encrypt(crypto1, key2)
	crypto3 := Encrypt(crypto2, key3)
	crypto4 := Encrypt(crypto3, key4)

	crypto5 := Decrypt(crypto4, key2)
	crypto6 := Decrypt(crypto5, key3)
	crypto7 := Decrypt(crypto6, key1)
	decrypted := Decrypt(crypto7, key4)

	fmt.Printf("%v\n", decrypted)

	if decrypted.String() != "1331" {
		t.Fail()
	}
}

func TestSmallerKeys(t *testing.T) {
	key1 := GenerateKey(256)
	message := big.NewInt(1331)
	crypto1 := Encrypt(message, key1)
	fmt.Printf("SMALL: %v\n\n", key1.Encryption)

	decrypted := Decrypt(crypto1, key1)

	fmt.Printf("%v\n", decrypted)

	if decrypted.String() != "1331" {
		t.Fail()
	}
}
