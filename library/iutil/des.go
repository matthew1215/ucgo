package iutil

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func DesDecrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	out := make([]byte, len(data))
	blockMode.CryptBlocks(out, data)
	out = PKCS5UnPadding(out)
	return out, nil
}

func DesEncrypt(data, key []byte, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	blockMode.CryptBlocks(out, data)
	return out, nil
}

//明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
