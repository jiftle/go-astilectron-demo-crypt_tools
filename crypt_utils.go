package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"strings"
)

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func encrypt_triple_des(src []byte, key []byte) []byte {
	//fmt.Printf("-->> src: %v\nkey: %v\n", src, key)
	block, _ := des.NewTripleDESCipher(key)
	//fmt.Printf("-->> block size: %v\n", block.BlockSize())
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	blockmode := cipher.NewCBCEncrypter(block, iv)
	blockmode.CryptBlocks(src, src)
	return src
}

func decrypt_triple_des(src []byte, key []byte) []byte {
	block, _ := des.NewTripleDESCipher(key)
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	blockmode := cipher.NewCBCDecrypter(block, iv)
	blockmode.CryptBlocks(src, src)
	//src = unpadding(src)
	return src
}
func HexStrToBytes(str string) []byte {
	b, _ := hex.DecodeString(str)
	return b
}
func BytesToHexStr(b []byte) string {
	h := fmt.Sprintf("%x", b)
	rh := strings.ToUpper(h)
	return rh
}
func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
