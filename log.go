package log

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/poolqa/log/rotator"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"sync"
)

//var LogFilePath = "./log.conf" // your can change the file path before init
var _logger *logging.Logger

var once sync.Once
var _cfg *LogConfig

//func init() {
//	initLogger(LogFilePath)
//}

func Default() {
	var err error
	_cfg, err = LoadLogConfigJson(GetDefaultLogConfig())
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
	once.Do(initLogger)
}

func InitByConfigFile(filePath string) {
	var err error
	_cfg, err = LoadLogConfigFile(filePath)
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
	once.Do(initLogger)
}

func InitByConfigJson(configJson string) {
	var err error
	_cfg, err = LoadLogConfigJson([]byte(configJson))
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
	once.Do(initLogger)
}

func InitByConfigStruct(conf *LogConfig) {
	_cfg = conf
	once.Do(initLogger)
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

func initLogger() {
	_logger = logging.MustGetLogger("")
	_logger.ExtraCalldepth = 1
	// check log dir path
	err := initLogDir(_cfg.FileName)
	if err != nil {
		log.Fatal(err)
	}

	logLevel := getLogLevel(_cfg.Level)
	//logging.SetLevel(logLevel, "")

	backendArr := []logging.Backend{}
	format := logging.MustStringFormatter(_cfg.Format)
	if _cfg.HasConsole {
		backendConsole := logging.NewLogBackend(os.Stdout, "", 0)
		backendConsole.Color = _cfg.Color
		backendConsoleFormatter := logging.NewBackendFormatter(backendConsole, format)
		backendConsoleLvLeveled := logging.AddModuleLevel(backendConsoleFormatter)
		backendConsoleLvLeveled.SetLevel(logLevel, "")
		backendArr = append(backendArr, backendConsoleLvLeveled)
	}
	rotateMode := getLogRotateMode(_cfg.DateSlice)
	lf := rotator.NewLogger(_cfg.FileName, _cfg.MaxSize, _cfg.MaxAge, rotateMode, false)
	backendNormal := logging.NewLogBackend(lf, "", 0)
	backendNormalFormatter := logging.NewBackendFormatter(backendNormal, format)
	backendNormalLvLeveled := logging.AddModuleLevel(backendNormalFormatter)
	backendNormalLvLeveled.SetLevel(logLevel, "")
	backendArr = append(backendArr, backendNormalLvLeveled)

	if _cfg.LevelFileName != nil && len(_cfg.LevelFileName) > 0 {
		for k, fp := range _cfg.LevelFileName {
			lvFileLevel := getLogLevel(k)
			lvF := rotator.NewLogger(fp, _cfg.MaxSize, _cfg.MaxAge, rotateMode, false)
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
