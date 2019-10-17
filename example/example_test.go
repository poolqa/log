package main

import (
	"github.com/poolqa/log"
	"testing"
	"time"
)

func testConfigStruct() {
	conf := log.LogConfig{
		Level:    "debug",
		FileName: "./logs/log.log",
		LevelFileName: map[string]string{
			"error": "./logs/error.log",
		},
		HasConsole: true,
		Color:      true,
		Json:       false,
		MaxSize:    0,
		MaxAge:     1,
		DateSlice:  "m",
		Format:     "%{time:2006/01/02 15:04:05.000} %{shortfile} [%{level:.4s}] %{message}",
	}
	log.InitByConfigStruct(&conf)
}

func Test(t *testing.T) {
	//log.Default()
	//log.InitByConfigFile("./log.conf")
	//log.InitByConfigJson(string(log.GetDefaultLogConfig()))
	testConfigStruct()
	log.Debugf("debug %s", "test")
	log.Info("info")
	log.Notice("notice")
	log.Warn("warning")
	log.Error("err")
	log.Critical("crit")
	log.Fatal("fatal")
	for {
		log.Error("age test......")
		time.Sleep(500 * time.Millisecond)
	}
}

func BenchmarkLogTextPositive(b *testing.B) {
	log.Default()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Error("The quick brown fox jumps over the lazy dog")
		}
	})
}
