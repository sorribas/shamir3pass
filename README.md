shamir3pass
===========

Implementation of the [shamir three pass protocol](https://en.wikipedia.org/wiki/Three-pass_protocol) encryption function in go.

DISCLAIMER: This package has not been thoroughly tested. Feedback about potential security flaws in it are welcome.

The shamir three pass protocol encryption function is commutative, which means that if two or more keys are generated
based on the same large prime number, encryption and decryption can be done in any order for the different keys. This
means that an exchange like the following one is possible.

1. Alice encrypts a message with key1 and sends it to Bob
2. Bob encrypts the already encrypted message with key2 and sends it back to Alice
3. Alice decrypts the message she received from Bob with key1 and sends it back to him.
4. Bob decrypts this last message with key2 revealing the original message.

This is why the protocol is called three pass protocol, because with three passes a message
can be securely sent without the need to exchange keys, given that the parts agree on a
large prime number to generate their keys.

Usage
-----

Basic usage. Encryption with one key.

```go
package main

import (
	"fmt"
	"github.com/sorribas/shamir3pass"
	"math/big"
)

func main() {
	message := []byte("Hello world!")

	key := shamir3pass.GenerateKey(nil) // Generate a key based on a random prime
	messageBigInt := &big.Int{}
	messageBigInt.SetBytes(message) // Converting the message bytes in to a big int

	ciphertext := shamir3pass.Encrypt(messageBigInt, key)
	decrypted := shamir3pass.Decrypt(ciphertext, key).Bytes()

	fmt.Printf("%v\n", string(decrypted)) // prints "Hello world!"
}
```

Encryption with two keys.

```go
package main

import (
	"fmt"
	"github.com/sorribas/shamir3pass"
	"math/big"
)

func main() {
	message := []byte("Hello world!")

	prime := shamir3pass.Random1024BitPrime()

	key1 := shamir3pass.GenerateKey(prime)
	key2 := shamir3pass.GenerateKey(prime)
	messageBigInt := &big.Int{}
	messageBigInt.SetBytes(message)

	ciphertext1 := shamir3pass.Encrypt(messageBigInt, key1)
	ciphertext2 := shamir3pass.Encrypt(ciphertext1, key2)
	ciphertext3 := shamir3pass.Decrypt(ciphertext2, key1)
	decrypted := shamir3pass.Decrypt(ciphertext3, key2).Bytes()

	fmt.Printf("%v\n", string(decrypted)) // prints "Hello world!"
}
```

To do
-----

* Allow configurable key size. It is now hardcoded to 2048 bits.
* Improve docs.

License
-------

MIT
