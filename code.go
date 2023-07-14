package gotools

// 编码操作

import (
	"io/ioutil"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

/**
 * @description: 跳过文件BOM头
 * @param {string} filename 文件路径
 * @return {newFile, err} 剔除BOM头的字节数组, 错误
 */
func SkipBOM(filename string) (newFile []byte, err error) {
	file, err := ioutil.ReadFile(filename)
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

// 转换为utf8编码
func Convert(data []byte) ([]byte, error) {
	if utf8.Valid(data) {
		return data, nil
	} else {
		utf8Data, err := simplifiedchinese.GBK.NewDecoder().Bytes(data)
		if err != nil {
			return utf8Data, err
		} else {
			return utf8Data, nil
		}
	}
}
