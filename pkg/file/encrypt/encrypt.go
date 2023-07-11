package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func EncryptFile(key []byte, inputFile string, outputFile string) error {
	plaintext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// 创建一个GCM模式的AES加密器
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// 生成一个随机的nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// 使用GCM加密模式进行加密
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// 将加密后的内容写入文件
	err = ioutil.WriteFile(outputFile, ciphertext, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Encryption completed successfully.")
	return nil
}
