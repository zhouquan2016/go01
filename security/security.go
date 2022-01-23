package security

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
)

// Md5 to encode string of src,
// the return string length is 32
func Md5(src string) string {
	bs := md5.Sum([]byte(src))
	endCodeStr := hex.EncodeToString(bs[:])
	return endCodeStr
}

func GenRsaKeys(bits int) (priKey, pubKey string, err error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	err = key.Validate()
	if err != nil {
		return "", "", err
	}
	priBytes := x509.MarshalPKCS1PrivateKey(key)
	priKey = string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: priBytes}))
	pubBytes := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	pubKey = string(pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: pubBytes}))
	return priKey, pubKey, err
}
