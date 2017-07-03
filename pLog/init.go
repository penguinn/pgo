package pLog

import (
	"fmt"
	log "github.com/cihub/seelog"
)

func Init(logFile string) {
	logger, err := log.LoggerFromConfigAsFile(logFile)
	if err != nil {
		fmt.Println("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)
}
