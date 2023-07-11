package main

import (
	"fmt"
	"github.com/YuanJey/goutils2/pkg/file/decrypt"
	"github.com/YuanJey/goutils2/pkg/file/encrypt"
	"os"
)

func main() {
	key := []byte("0123456789abcdef") // 16字节的AES密钥

	inputFilepath := "D:\\TEST.png"
	encryptedFilepath := "D:\\encrypted.png"
	decryptedFilepath := "D:\\decrypted.png"

	// 加密文件
	err := encrypt.EncryptFile(key, inputFilepath, encryptedFilepath)
	if err != nil {
		fmt.Println("Encryption error:", err)
		os.Exit(1)
	}

	// 解密文件
	err = decrypt.DecryptFile(key, encryptedFilepath, decryptedFilepath)
	if err != nil {
		fmt.Println("Decryption error:", err)
		os.Exit(1)
	}
}
