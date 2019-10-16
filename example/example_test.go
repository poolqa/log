package main

import (
	"github.com/poolqa/log"
	"testing"
)

func Test(t *testing.T) {
	log.Default()
	log.InitByConfigFile("./log.conf")
	log.InitByConfigJson(string(log.GetDefaultLogConfig()))
	log.Debugf("debug %s", "test")
	log.Info("info")
	log.Warn("warning")
	log.Error("err")
	log.Critical("crit")
	log.Fatal("fatal")
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
