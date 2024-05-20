package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func NewRSA() *RSA {
	return &RSA{}
}

type RSA struct {
}

func (ins *RSA) GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return key, &key.PublicKey, nil
}

func (ins *RSA) EncodePKCS1(key *rsa.PrivateKey) (privateKey, publicKey []byte) {
	p1 := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	p2 := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
	})
	return p1, p2
}

func (ins *RSA) EncodePKCS8(key *rsa.PrivateKey) (privateKey, publicKey []byte, err error) {
	b1, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, nil, err
	}
	b2, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: b1}),
		pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: b2}), nil
}

func (ins *RSA) GenerateRandomArt(key *rsa.PublicKey) (string, string) {
	b2, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", ""
	}
	hash := sha256.Sum256(b2)
	return base64.StdEncoding.EncodeToString(hash[:]), generateRandomArt(b2)
}
