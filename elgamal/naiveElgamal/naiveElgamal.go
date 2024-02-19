package naiveElgamal

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const P = 23

var g = big.NewInt(5)

func GenerateKey() (*big.Int, *big.Int) {
	publicKey, _ := rand.Int(rand.Reader, big.NewInt(P-2))
	publicKey = publicKey.Add(publicKey, big.NewInt(2))

	privateKey := new(big.Int).Exp(g, publicKey, big.NewInt(P))

	return publicKey, privateKey
}

func Encrypt(m, publicKey *big.Int) (*big.Int, *big.Int) {

	r, _ := rand.Int(rand.Reader, big.NewInt(P-2))
	r = r.Add(r, big.NewInt(2))

	c1 := new(big.Int).Exp(g, r, big.NewInt(P))
	c2 := new(big.Int).Mul(m, new(big.Int).Exp(publicKey, r, big.NewInt(P)))
	c2 = c2.Mod(c2, big.NewInt(P))

	return c1, c2
}

func Decrypt(privateKey, c1, c2 *big.Int) *big.Int {

	m := new(big.Int).Mul(c2, new(big.Int).Exp(c1, new(big.Int).Sub(big.NewInt(P-1), privateKey), big.NewInt(P)))
	m = m.Mod(m, big.NewInt(P))
	return m
}

func test() {
	privateKey, publicKey := GenerateKey()

	m := big.NewInt(16)

	c1, c2 := Encrypt(m, publicKey)
	fmt.Printf("Enc: (%s, %s)\n", c1, c2)

	mDecrypted := Decrypt(privateKey, c1, c2)
	fmt.Printf("Dec: %s\n", mDecrypted)
}
