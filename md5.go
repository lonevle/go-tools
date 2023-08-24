package gotools

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

/**
 * @description: 字符串md5计算
 * @param {string} str 需要计算的字符串
 * @return {*} md5字符串
 */
func MD5(str string) string { // md5
	sum := md5.Sum([]byte(str))
	// md5str := fmt.Sprintf("%x", has) //将[16]byte转成16进制
	return hex.EncodeToString(sum[:])
}

// 获取文件的md5
func GetFileMd5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5h := md5.New()
	_, err = io.Copy(md5h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5h.Sum(nil)), nil
}
