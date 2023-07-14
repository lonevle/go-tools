package gotools

import (
	"crypto/md5"
	"fmt"
)

/**
 * @description: 字符串md5计算
 * @param {string} str 需要计算的字符串
 * @return {*} md5字符串
 */
func MD5(str string) string { // md5
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
