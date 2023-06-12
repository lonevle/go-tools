package gotools

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/adhocore/jsonc"
)

/**
 * @description: 跳过文件BOM头
 * @param {string} filename 文件路径
 * @return {newFile, err} 剔除BOM头的字节数组, 错误
 */
func SkipBOM(filepath string) (newFile []byte, err error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	if len(file) >= 4 && isUTF32BigEndianBOM4(file) {

		return file[4:], nil
	}
	if len(file) >= 4 && isUTF32LittleEndianBOM4(file) {
		return file[4:], nil
	}
	if len(file) > 2 && isUTF8BOM3(file) {
		return file[3:], nil
	}
	if len(file) == 2 && isUTF16BigEndianBOM2(file) {
		return file[2:], nil
	}
	if len(file) == 2 && isUTF16LittleEndianBOM2(file) {
		return file[2:], nil
	}
	return file, nil
}

func isUTF32BigEndianBOM4(buf []byte) bool {
	if len(buf) < 4 {
		return false
	}
	return buf[0] == 0x00 && buf[1] == 0x00 && buf[2] == 0xFE && buf[3] == 0xFF
}

func isUTF32LittleEndianBOM4(buf []byte) bool {
	if len(buf) < 4 {
		return false
	}
	return buf[0] == 0xFF && buf[1] == 0xFE && buf[2] == 0x00 && buf[3] == 0x00
}

func isUTF8BOM3(buf []byte) bool {
	if len(buf) < 3 {
		return false
	}
	return buf[0] == 0xEF && buf[1] == 0xBB && buf[2] == 0xBF
}

func isUTF16BigEndianBOM2(buf []byte) bool {
	if len(buf) < 2 {
		return false
	}
	return buf[0] == 0xFE && buf[1] == 0xFF
}

func isUTF16LittleEndianBOM2(buf []byte) bool {
	if len(buf) < 2 {
		return false
	}
	return buf[0] == 0xFF && buf[1] == 0xFE
}

/**
 * @description: 格式化不规则json
 * @param {string} configFile json文件路径
 * @return {[]byte} json文件字节数组
 */
func JsonStrip(configFile string) []byte {
	configByte, err := SkipBOM(configFile)
	if err != nil {
		log.Fatalln(err)
		// panic(err)
	}
	j := jsonc.New()
	return j.Strip(configByte)
}

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
