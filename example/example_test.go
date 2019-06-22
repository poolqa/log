package main

import (
	"../../log"
	"testing"
)

func Test(t *testing.T) {
	log.Debugf("debug %s", "test")
	log.Info("info")
	log.Notice("notice")
	log.Warn("warning")
	log.Error("err")
	log.Critical("crit")
	log.Fatal("fatal")
	for i := 10000; i > 0; i-- {
		log.Debugf("debug %s", "test")
		log.Info("info")
		log.Notice("notice")
		log.Warn("warning")
		log.Error("err")
		log.Critical("crit")
		log.Fatal("fatal")
	}
}

func BenchmarkLogTextPositive(b *testing.B) {

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Error("The quick brown fox jumps over the lazy dog")
		}
	})
}