package log

import (
	"fmt"
	"github.com/poolqa/log/rotator"
	"github.com/rs/zerolog"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

//var LogFilePath = "./log.conf" // your can change the file path before init
var _logger *zerolog.Logger

var once sync.Once
var _cfg *logConfig

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

type timeHook struct{}

func (h timeHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("time", time.Now().Format("2006/01/02 15:04:05.000"))
}

func LevelFormatter(i interface{}) string {
	return "[" + strings.ToUpper(i.(string)[:4]) + "]"
	//return i.(string)
}

func TimeFormatter(i interface{}) string {
	return i.(string)
}

func MessageFormatter(i interface{}) string {
	return i.(string)
}

func initLogger() {
	// check log dir path
	err := initLogDir(_cfg.FileName)
	if err != nil {
		log.Fatal(err)
	}
	writer := []io.Writer{}
	if _cfg.HasConsole {
		output := zerolog.ConsoleWriter{Out: os.Stdout}
		output.FormatLevel = LevelFormatter
		output.FormatTimestamp = TimeFormatter
		output.FormatMessage = MessageFormatter
		writer = append(writer, output)
	}
	rotateMode := getLogRotateMode(_cfg.DateSlice)
	lf := rotator.NewLogger(_cfg.FileName, _cfg.MaxSize, _cfg.MaxAge, rotateMode, false)
	output := zerolog.ConsoleWriter{Out: lf}
	output.FormatLevel = LevelFormatter
	output.FormatTimestamp = TimeFormatter
	output.FormatMessage = MessageFormatter
	writer = append(writer, output)

	//if _cfg.LevelFileName != nil && len(_cfg.LevelFileName) > 0 {
	//	for k, fp := range _cfg.LevelFileName {
	//		lvFileLevel := getLogLevel(k)
	//		lvF := rotator.NewLogger(fp, _cfg.MaxSize, _cfg.MaxAge, rotateMode, false)
	//		backendLv := logging.NewLogBackend(lvF, "", 0)
	//		backendLvFormatter := logging.NewBackendFormatter(backendLv, format)
	//		backendLvLeveled := logging.AddModuleLevel(backendLvFormatter)
	//		backendLvLeveled.SetLevel(lvFileLevel, "")
	//		backendArr = append(backendArr, backendLvLeveled)
	//	}
	//}

	// Set the backends to be used.
	outs := io.MultiWriter(writer...)
	logger := zerolog.New(outs).With().Logger().Hook(timeHook{}) // .Level()
	//_logger.ExtraCalldepth = 1
	_logger = &logger

}

//log critical level
func Fatal(a ...interface{}) {
	_logger.Fatal().Msg(fmt.Sprint(a...))
}

//log critical format
func Fatalf(format string, a ...interface{}) {
	_logger.Fatal().Msgf(format, a...)
}

//log error level
func Error(a ...interface{}) {
	_logger.Error().Msg(fmt.Sprint(a...))
}

//log error format
func Errorf(format string, a ...interface{}) {
	_logger.Error().Msgf(format, a...)
}

//log warning level
func Warn(a ...interface{}) {
	_logger.Warn().Msg(fmt.Sprint(a...))
}

//log warning format
func Warnf(format string, a ...interface{}) {
	_logger.Warn().Msgf(format, a...)
}

//log info level
func Info(a ...interface{}) {
	_logger.Info().Msg(fmt.Sprint(a...))
}

//log info format
func Infof(format string, a ...interface{}) {
	_logger.Info().Msgf(format, a...)
}

//log debug level
func Debug(a ...interface{}) {
	_logger.Debug().Msg(fmt.Sprint(a...))
}

//log debug format
func Debugf(format string, a ...interface{}) {
	_logger.Debug().Msgf(format, a...)
}
