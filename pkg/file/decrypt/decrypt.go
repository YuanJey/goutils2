package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func DecryptFile(key []byte, inputFile string, outputFile string) error {
	ciphertext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// 创建一个GCM模式的AES解密器
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// 从密文中提取nonce
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return fmt.Errorf("Invalid ciphertext")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// 使用GCM解密模式进行解密
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// 将解密后的内容写入文件
	err = ioutil.WriteFile(outputFile, plaintext, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Decryption completed successfully.")
	return nil
}
