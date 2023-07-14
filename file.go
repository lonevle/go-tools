package gotools

// 文件相关操作

import (
	"os"

	"github.com/adhocore/jsonc"
)

/**
 * @description: 检查路径是否存在
 * @param {string} path 路径
 * @return {bool, error} 是否存在, error
 */
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 读取文本文件
func ReadFile(File string) ([]byte, error) {
	fileByte, err := SkipBOM(File)
	if err != nil {
		return fileByte, err
	}
	utf8Byte, err := Convert(fileByte)
	if err != nil {
		return utf8Byte, err
	}
	return utf8Byte, nil
}

// 读取json文件
func ReadJson(jsonFile string) ([]byte, error) {
	fileByte, err := SkipBOM(jsonFile)
	if err != nil {
		return fileByte, err
	}
	utf8Byte, err := Convert(fileByte)
	if err != nil {
		return utf8Byte, err
	}
	j := jsonc.New()
	return j.Strip(utf8Byte), nil
}
