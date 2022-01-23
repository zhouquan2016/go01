package security

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

const (
	privateBlockType = "RSA PRIVATE KEY"
	publicBlockType  = "RSA PUBLIC KEY"
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
	priKey = string(pem.EncodeToMemory(&pem.Block{Type: privateBlockType, Bytes: priBytes}))
	pubBytes := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	pubKey = string(pem.EncodeToMemory(&pem.Block{Type: publicBlockType, Bytes: pubBytes}))
	return priKey, pubKey, err
}

func SignWithRsa(privateKey string, src string) (string, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil || block.Type != privateBlockType {
		return "", errors.New("failed to decode PEM block containing private key")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	hash := crypto.SHA256.New()
	hash.Write([]byte(src))
	hashed := hash.Sum(nil)
	signBytes, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signBytes), nil
}

func VerifyWithRsa(signData string, srcData string, publicKey string) (bool, error) {

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil || block.Type != publicBlockType {
		return false, errors.New("failed to decode PEM block containing private key")
	}
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	bytes, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return false, err
	}
	hash := crypto.SHA256.New()
	hash.Write([]byte(srcData))
	hashed := hash.Sum(nil)
	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, hashed, bytes)
	if err != nil {
		return false, err
	}
	return true, err
}
