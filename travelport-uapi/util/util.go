package util

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/redochen/demos/travelport-uapi/config"
	CcFile "github.com/redochen/tools/file"
	Cchttp "github.com/redochen/tools/http"
)

//PostRequest 发送请求到接口
func PostRequest(pcc *config.PCC, svcName string, buf []byte) (string, error) {
	if nil == pcc {
		return "", errors.New("PCC is nil")
	}

	req := &Cchttp.Request{
		URL:            pcc.EndPoint + svcName,
		Method:         "POST",
		AcceptEncoding: "gzip,deflate",
		ContentType:    "text/xml;charset=UTF-8",
		UserName:       pcc.UserID,
		Password:       pcc.Password,
		PostString:     string(buf),
		TimeoutSeconds: 60,
	}

	return req.Post()
}

//GetCabinClass 获取舱位等级：Premium First, First, Business, Economy, Premium Economy
func GetCabinClass(cabin string) string {
	cabinClass := "Economy"

	switch strings.ToUpper(cabin) {
	case "P":
		cabinClass = "PremiumFirst"
		break
	case "F":
		cabinClass = "First"
		break
	case "C":
		cabinClass = "Business"
		break
	case "S":
		cabinClass = "PremiumEconomy"
	}

	return cabinClass
}

//DumpFile 保存DUMP文件
func DumpFile(file string, val string, force bool) {
	if !config.OutputToFile && !force {
		return
	}

	dir := "dump"
	_, err := os.Stat(dir)
	if err != nil {
		os.Mkdir(dir, os.ModeDir)
	}

	now := time.Now()
	folder := now.Format("2006-01-02")
	path := fmt.Sprintf("%s/%s", dir, folder)
	_, err = os.Stat(path)
	if err != nil {
		os.Mkdir(path, os.ModeDir)
	}

	fullPath := path + "/" + file
	_, err = os.Stat(fullPath)
	if nil == err {
		os.Remove(fullPath)
	}

	go CcFile.DumpFile(fullPath, val)
}

//LoadFile 加载文件
func LoadFile(file string) (string, error) {
	fe, err := CcFile.Open(file, false, true)
	if err != nil {
		return "", err
	}

	defer fe.Close()
	len, err := fe.Size()

	if err != nil {
		return "", err
	} else if len == 0 {
		return "", errors.New("lenght is zero")
	}

	data := make([]byte, len)
	_, err = fe.ReadEx(data, 0, false)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
