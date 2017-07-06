package pLog

import (
	"fmt"
	log "github.com/cihub/seelog"
	"errors"
)

func Init(logFile string) {
	if len(logFile) == 0{
		panic(errors.New("empty logFile"))
	}
	logger, err := log.LoggerFromConfigAsFile(logFile)
	if err != nil {
		fmt.Println("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)
}
