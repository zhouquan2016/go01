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
	"log"
	"time"
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
	startTime := time.Now()
	defer func() {
		log.Println("生成rsa公私钥耗时:", time.Now().Sub(startTime).String())
	}()
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

func decodePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil || block.Type != privateBlockType {
		return nil, errors.New("failed to decode PEM block containing private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

//私有加签
func SignWithPrivateKey(privateKey string, src []byte) (string, error) {
	startTime := time.Now()
	defer func() {
		log.Println("rsa生成签名耗时:", time.Now().Sub(startTime).String())
	}()
	key, err := decodePrivateKey(privateKey)
	if err != nil {
		return "", err
	}
	hash := crypto.SHA256.New()
	hash.Write(src)
	hashed := hash.Sum(nil)
	signBytes, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signBytes), nil
}

func decodePublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil || block.Type != publicBlockType {
		return nil, errors.New("failed to decode PEM block containing private key")
	}
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

// VerifyWithPublicKey 公钥验签
func VerifyWithPublicKey(signData string, srcData []byte, publicKey string) (bool, error) {
	startTime := time.Now()
	defer func() {
		log.Println("rsa验证签名耗时:", time.Now().Sub(startTime).String())
	}()
	key, err := decodePublicKey(publicKey)
	if err != nil {
		return false, err
	}

	bytes, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return false, err
	}
	hash := crypto.SHA256.New()
	hash.Write(srcData)
	hashed := hash.Sum(nil)
	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, hashed, bytes)
	if err != nil {
		return false, err
	}
	return true, err
}

//  DecryptWithPrivateKey私钥解密
func DecryptWithPrivateKey(privateKey string, encryptData string) ([]byte, error) {
	startTime := time.Now()
	defer func() {
		log.Println("rsa解密耗时:", time.Now().Sub(startTime).String())
	}()
	key, err := decodePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	encryptBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, key, encryptBytes)
}

//EncryptWithPublicKey 公钥加密
func EncryptWithPublicKey(bs []byte, publicKey string) (string, error) {
	key, err := decodePublicKey(publicKey)
	if err != nil {
		return "", err
	}
	encryptBytes, err := rsa.EncryptPKCS1v15(rand.Reader, key, bs)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptBytes), nil

}
