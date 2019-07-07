package log

import (
	"bytes"
	"encoding/json"
	"github.com/op/go-logging"
	"github.com/poolqa/log/rotator"
	"os"
	"strings"
)

const gDefaultConfigJson = `
# 備註說明：備註不能寫在後面，只能單行由#開頭
# Level : debug, info, notice, warn, error, critical(fatal)
#	MaxSize : size * mb
# Color : 只有在console才會生效
# DateSlice : d:天, h:小時
# Format :
# %{id}        Sequence number for log message (uint64).
# %{pid}       Process id (int)
# %{time}      Time when log occurred (time.Time) ex:%{time:2006/01/02 15:04:05.000}
# %{level}     Log level (Level)
# %{module}    Module (string)
# %{program}   Basename of os.Args[0] (string)
# %{message}   Message (string)
# %{longfile}  Full file name and line number: /a/b/c/d.go:23
# %{shortfile} Final file name element and line number: d.go:23
# %{callpath}  Callpath like main.a.b.c...c  "..." meaning recursive call ~. meaning truncated path
# %{color}     ANSI color based on log level 不知道什麼時候才有效果
{
	"Level": "DEBUG",
	"FileName":"./logs/log.log",
	"LevelFileName": {
#		"error": "./logs/error.log"
	},
	"HasConsole": true,
	"Color": true,
	"MaxSize": 0,
	"DateSlice": "d",
	"Format": "%{time:2006/01/02 15:04:05.000} %{shortfile} [%{level:.4s}] %{message}"
}
`

type logConfig struct {
	Level         string
	FileName      string
	LevelFileName map[string]string
	HasConsole    bool
	Color         bool
	Json          bool
	MaxSize       int
	MaxAge        int
	DateSlice     string
	Format        string
}

func LoadLogConfig(configJsonFile string) (*logConfig, error) {
	fp, err := os.Open(configJsonFile)
	if err == nil {
		defer fp.Close()

		fileInfo, err := fp.Stat()
		if err != nil {
			return nil, err
		}
		buffer := make([]byte, fileInfo.Size())
		_, err = fp.Read(buffer) // 文件内容读取到buffer中
		if err != nil {
			return nil, err
		}
		buffer = removeConfRemark(buffer)
		config := logConfig{}
		err = json.Unmarshal([]byte(buffer), &config)
		return &config, err
	} else {
		// get default
		buffer := removeConfRemark([]byte(gDefaultConfigJson))
		config := logConfig{}
		err = json.Unmarshal(buffer, &config)
		return &config, err
	}
}
func removeConfRemark(bConf []byte) []byte {
	sConfLines := strings.Split(string(bConf), "\n")
	buffer := bytes.Buffer{}
	for _, line := range sConfLines {
		newLine := strings.TrimSpace(line)
		if len(newLine) == 0 || newLine[0] == '#' {
			continue
		}
		buffer.WriteString(line + "\n")
	}
	out := make([]byte, buffer.Len())
	_, _ = buffer.Read(out)
	return out
}

func getLogLevel(strLv string) logging.Level {
	switch strings.ToLower(strLv) {
	case "debug":
		return logging.DEBUG
	case "info":
		return logging.INFO
	case "notice":
		return logging.NOTICE
	case "warn":
		return logging.WARNING
	case "error":
		return logging.ERROR
	case "fatal", "critical":
		return logging.CRITICAL
	default:
		return logging.DEBUG
	}
}

func getLogRotateMode(strMode string) rotator.RotateDateMode {
	switch strings.ToLower(strMode) {
	case "d":
		return rotator.ROTATE_DATE_MODE_DAY
	case "h":
		return rotator.ROTATE_DATE_MODE_HOUR
	default:
		return rotator.ROTATE_DATE_MODE_NO
	}
}
