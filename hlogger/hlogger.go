package hlogger

import (
	"log"
	"os"
	"sync"
)

type hydraLogger struct {
	*log.Logger
	filename string
}

var hlogger *hydraLogger
var once sync.Once

func GetInstance() *hydraLogger {
	once.Do(func() {
		hlogger = createLogger("logger.log")
	})
	return hlogger
}

func createLogger(filename string) *hydraLogger {
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &hydraLogger{
		filename: filename,
		Logger:   log.New(file, "Hydra ", log.Lshortfile),
	}
}
