package log

import (
	"./rotator"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/op/go-logging"
	"os"
	"strings"
)

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
	if err != nil {
		fmt.Printf("err:%v", err)
		return nil, err
	}
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
	//strConf, _ := GoJsoner.Discard(string(buffer))
	err = json.Unmarshal([]byte(buffer), &config)
	return &config, err
}
func removeConfRemark(bConf []byte) []byte {
	sConfLines := strings.Split(string(bConf), "\r\n")
	buffer := bytes.Buffer{}
	for _, line := range sConfLines {
		newLine := strings.TrimLeft(line, " \t\r\n")
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
