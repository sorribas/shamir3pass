package shamir3pass

import (
	"fmt"
	"math/big"
	"testing"
)

func TestBasic(t *testing.T) {
	key := GenerateKey(nil)
	message := big.NewInt(1331)

	crypto := Encrypt(message, key)
	decrypted := Decrypt(crypto, key)
	fmt.Printf("%v\n", decrypted)

	if decrypted.String() != "1331" {
		t.Fail()
	}
}

func TestDoubleEncryption(t *testing.T) {
	key := GenerateKey(nil)
	key2 := GenerateKey(key.Prime)

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
	key1 := GenerateKey(nil)
	key2 := GenerateKey(key1.Prime)
	key3 := GenerateKey(key1.Prime)
	key4 := GenerateKey(key1.Prime)

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
