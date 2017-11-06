//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"

	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
)

type aesEncryption struct {
	key string
	iv  string
}

func NewAESEncryption(encData export.EncryptionDetails) Transformer {
	aesData := aesEncryption{
		key: encData.Key,
		iv:  encData.InitVector,
	}
	return aesData
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (aesData aesEncryption) Transform(data []byte) []byte {
	iv := make([]byte, 16)
	copy(iv, []byte(aesData.iv))

	hash := sha1.New()

	hash.Write([]byte((aesData.key)))
	key := hash.Sum(nil)
	key = key[:16]

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error("Error", zap.Error(err))
		return nil
	}

	ecb := cipher.NewCBCEncrypter(block, iv)
	content := pkcs5Padding(data, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	encodedData := []byte(base64.StdEncoding.EncodeToString(crypted))

	return encodedData
}
