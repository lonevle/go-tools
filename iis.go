package gotools

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

// StartPool 启动应用程序池
func StartPool(appPoolName string) error {
	cmd := exec.Command(`C:\Windows\System32\inetsrv\appcmd.exe`, "start", "apppool", fmt.Sprintf("/apppool.name:%s", appPoolName))

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := ConvertByte2String(stdout.Bytes(), GB18030), string(stderr.String())

	log.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		return err
	}
	return nil

}

// StopPool 停止应用程序池
func StopPool(appPoolName string) error {
	cmd := exec.Command(`C:\Windows\System32\inetsrv\appcmd.exe`, "stop", "apppool", fmt.Sprintf("/apppool.name:%s", appPoolName))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := ConvertByte2String(stdout.Bytes(), GB18030), string(stderr.String())

	log.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		if strings.Contains(outStr, "已停止") {
			return nil
		}
		return err
	}
	return nil
}

// StartSite 启动网站
func StartSite(appSiteName string) error {
	cmd := exec.Command(`C:\Windows\System32\inetsrv\appcmd.exe`, "start", "site", appSiteName)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := ConvertByte2String(stdout.Bytes(), GB18030), string(stderr.String())

	log.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)

		return err
	}
	return nil
}

// StopPool 停止网站
func StopSite(appSiteName string) error {
	// log.Println("停止网站", fmt.Sprintf(`C:\Windows\System32\inetsrv\appcmd.exe stop site "%s"`, appSiteName))
	cmd := exec.Command(`C:\Windows\System32\inetsrv\appcmd.exe`, "stop", "site", appSiteName)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := ConvertByte2String(stdout.Bytes(), GB18030), string(stderr.String())

	log.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		return err
	}
	return nil
}
