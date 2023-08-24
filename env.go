package gotools

// 运行环境

import (
	"log"
	"os"
	"path/filepath"
)

// GetProgramPath 获取程序所在目录
// @return string 文件夹路径
func GetProgramPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(exePath)
}

// 获取程序名
func GetProgramName() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Base(exePath)
}
