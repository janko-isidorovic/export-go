//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"testing"

	"github.com/drasko/edgex-export"
)

const (
	plainString = "This is the test string used for testing"
	iv          = "123456789012345678901234567890"
	key         = "aquqweoruqwpeoruqwpoeruqwpoierupqoweiurpoqwiuerpqowieurqpowieurpoqiweuroipwqure"
)

func aesDecrypt(crypt []byte, key []byte) []byte {
	hash := sha1.New()

	hash.Write([]byte((key)))
	key = hash.Sum(nil)
	key = key[:16]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("key error")
	}

	ecb := cipher.NewCBCDecrypter(block, []byte(iv[:16]))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	trimmed := pkcs5Trimming(decrypted)
	decodedData, _ := base64.StdEncoding.DecodeString(string(trimmed))

	return decodedData
}

func pkcs5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func TestAES(t *testing.T) {

	aesData := export.EncryptionDetails{
		Algo:       "AES",
		Key:        key,
		InitVector: iv,
	}

	enc := NewAESEncryption(aesData)

	cphrd := enc.Transform([]byte(plainString))

	decphrd := aesDecrypt(cphrd, []byte(aesData.Key))

	if string(plainString) != string(decphrd) {
		t.Fatal("Encoded string ", cphrd, " is not ", decphrd)
	}
}
