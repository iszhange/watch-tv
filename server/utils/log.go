/*
日志
*/
package utils

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Log struct {
	date string
	log  *log.Logger
}

var inst Log

func init() {
	inst := Log{}
	inst.Init()
}

// 记录日志
// level: INFO DEBUG WARN ERROR
func LogMessage(level string, message string) {
	inst.Init()
	inst.log.Printf("[%s] %s\n", strings.ToUpper(level), message)
}

func InfoMessage(message string) {
	LogMessage("INFO", message)
}

func DebugMessage(message string) {
	LogMessage("DEBUG", message)
}

func WarnMessage(message string) {
	LogMessage("WARN", message)
}

func ErrorMessage(message string) {
	LogMessage("ERROR", message)
}

func (o *Log) Init() {
	// 判断是否需要重新初始化
	curDate := time.Now().Format("20060102")
	if curDate == o.date {
		return
	}
	o.date = curDate

	logFileName := "log_" + curDate + ".log"
	exePath, _ := os.Getwd()
	ps := string(os.PathSeparator)
	// 创建目录
	logPath := exePath + ps + "logs"
	_, err := os.Stat(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(logPath, 0755)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	// 打开文件句柄
	filePath := logPath + ps + logFileName
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, f)
	o.log = log.New(multiWriter, "", log.Ldate|log.Ltime)
}
