package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

/**
 * @description: AES加密文件
 * @param {string} path 文件路径
 * @param {[]byte} key 密钥
 * @return {*} 错误
 */
func AESEncryptFile(path string, key []byte) error {
	// 打开输入文件
	inputFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inputFile.Close() // 延迟关闭输入文件

	// 创建输出文件
	outputFile, err := os.OpenFile(path+".enc", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close() // 延迟关闭输出文件

	// 加密密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// 生成随机的IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return err
	}
	// 把IV写入输出文件
	if _, err := outputFile.Write(iv); err != nil {
		return err
	}

	// 加密流
	stream := cipher.NewCTR(block, iv)

	// 复制输入文件到输出文件，并加密
	if _, err := io.Copy(outputFile, &cipher.StreamReader{S: stream, R: inputFile}); err != nil {
		return err
	}

	return nil
}

/**
 * @description: AES解密文件
 * @param {string} path 文件路径
 * @param {[]byte} key 密钥
 * @return {*} 错误
 */
func AESDecryptFile(path string, key []byte) error {
	// 打开输入文件
	inputFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inputFile.Close() // 延迟关闭输入文件

	// 创建输出文件
	outputFile, err := os.OpenFile(path+".dec", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close() // 延迟关闭输出文件

	// 解密密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// 读取IV从输入文件
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(inputFile, iv); err != nil {
		return err
	}

	// 解密流
	stream := cipher.NewCTR(block, iv)

	// 复制输入文件到输出文件，并解密
	if _, err := io.Copy(outputFile, &cipher.StreamReader{S: stream, R: inputFile}); err != nil {
		return err
	}

	return nil
}
