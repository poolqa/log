package log

import (
	"./rotator"
	"fmt"
	"github.com/op/go-logging"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
)

var LogFilePath = "./log.conf" // your can change the file path before init
var _logger *logging.Logger

func init() {
	initLogger(LogFilePath)
}

func initLogDir(logFile string) error {
	dir := filepath.Dir(logFile)
	_, err := os.Stat(dir) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

func initLogger(confPath string) {
	cfg, err := LoadLogConfig(confPath)
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
	_logger = logging.MustGetLogger("")
	_logger.ExtraCalldepth = 1
	// check log dir path
	err = initLogDir(cfg.FileName)
	if err != nil {
		log.Fatal(err)
	}

	logLevel := getLogLevel(cfg.Level)
	logging.SetLevel(logLevel, "")

	backendArr := []logging.Backend{}
	format := logging.MustStringFormatter(cfg.Format)
	if cfg.HasConsole {
		backendConsole := logging.NewLogBackend(os.Stdout, "", 0)
		backendConsole.Color = cfg.Color
		backendConsoleFormatter := logging.NewBackendFormatter(backendConsole, format)
		backendArr = append(backendArr, backendConsoleFormatter)
	}
	rotateMode := getLogRotateMode(cfg.DateSlice)
	lf := rotator.NewLogger(cfg.FileName, cfg.MaxSize, cfg.MaxAge, rotateMode, false)
	backendNormal := logging.NewLogBackend(lf, "", 0)
	backendNormalFormatter := logging.NewBackendFormatter(backendNormal, format)
	backendArr = append(backendArr, backendNormalFormatter)

	if cfg.LevelFileName != nil && len(cfg.LevelFileName) > 0 {
		for k, fp := range cfg.LevelFileName {
			lvFileLevel := getLogLevel(k)
			lvF := rotator.NewLogger(fp, cfg.MaxSize, cfg.MaxAge, rotateMode, false)
			backendLv := logging.NewLogBackend(lvF, "", 0)
			backendLvFormatter := logging.NewBackendFormatter(backendLv, format)
			backendLvLeveled := logging.AddModuleLevel(backendLvFormatter)
			backendLvLeveled.SetLevel(lvFileLevel, "")
			backendArr = append(backendArr, backendLvLeveled)
		}
	}
	// Set the backends to be used.
	logging.SetBackend(backendArr...)

}

//log critical level
func Critical(a ...interface{}) {
	_logger.Critical(a...)
}

//log critical format
func Criticalf(format string, a ...interface{}) {
	_logger.Criticalf(format, a...)
}

//log critical level
func Fatal(a ...interface{}) {
	_logger.Critical(a...)
}

//log critical format
func Fatalf(format string, a ...interface{}) {
	_logger.Criticalf(format, a...)
}

//log error level
func Error(a ...interface{}) {
	_logger.Error(a...)
}

//log error format
func Errorf(format string, a ...interface{}) {
	_logger.Errorf(format, a...)
}

//log warning level
func Warn(a ...interface{}) {
	_logger.Warning(a...)
}

//log warning format
func Warnf(format string, a ...interface{}) {
	_logger.Warningf(format, a...)
}

//log notice level
func Notice(a ...interface{}) {
	_logger.Notice(a...)
}

//log notice format
func Noticef(format string, a ...interface{}) {
	_logger.Noticef(format, a...)
}

//log info level
func Info(a ...interface{}) {
	_logger.Info(a...)
}

//log info format
func Infof(format string, a ...interface{}) {
	_logger.Infof(format, a...)
}

//log debug level
func Debug(a ...interface{}) {
	_logger.Debug(a...)
}

//log debug format
func Debugf(format string, a ...interface{}) {
	_logger.Debugf(format, a...)
}

func printError(message string) {
	fmt.Println(message)
	os.Exit(0)
}
