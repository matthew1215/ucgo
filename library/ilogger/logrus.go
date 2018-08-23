package ilogger

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"time"
	"ucgo/configs"
	"ucgo/library/iutil"
)

var logger *logrus.Logger

func getLogger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
		logger.Formatter = &logrus.JSONFormatter{}

		//logClient.SetLevel(logrus.DebugLevel)
		//apiLogPath := "api.log"
		//logWriter, err := rotatelogs.New(
		//	apiLogPath+".%Y-%m-%d-%H-%M.log",
		//	rotatelogs.WithLinkName(apiLogPath), // 生成软链，指向最新日志文件
		//	rotatelogs.WithMaxAge(7*24*time.Hour), // 文件最大保存时间
		//	rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		//)
		//writeMap := lfshook.WriterMap{
		//	logrus.InfoLevel:  logWriter,
		//	logrus.FatalLevel: logWriter,
		//}
		//lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
		//logClient.AddHook(lfHook)
	}
	filename := configs.GetConf().Logpath + time.Now().Format("2006-01-02")
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	logger.Out = file
	return logger
}

func getLoggerEntry() *logrus.Entry {
	currentDir, _ := os.Getwd()
	_, filename, line, _ := runtime.Caller(2)
	file := filename[len(currentDir):]
	return getLogger().WithFields(logrus.Fields{
		"guid":     iutil.GetGuid(),
		"deviceid": iutil.GetDeviceid(),
		"path":     iutil.GetPath(),
		"file":     file,
		"line":     line,
	})
}

func Info(v ...interface{}) {
	getLoggerEntry().Info(v)
}

func Warn(v ...interface{}) {
	getLoggerEntry().Warn(v)
}

func Fatal(v ...interface{}) {
	getLoggerEntry().Fatal(v)
}
