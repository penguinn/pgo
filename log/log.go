package pLog

import (
	"fmt"
	log "github.com/cihub/seelog"
	"errors"
	"github.com/penguinn/pgo/utils"
)

var logFileString = `
<seelog>
    <outputs formatid="main">
        <filter levels="critical,error,warn,info,debug">
            <buffered size="10000" flushperiod="1000">
                <rollingfile type="size" filename="log/grape.log" maxsize="10000000" maxrolls="30"/><!---向文件输出-->
                <!---
                <rollingfile type="date" filename="log/iat_server.log" datepattern="2006.01.02" maxrolls="30"/>
                -->
            </buffered>
        </filter>
        <console />  <!---向屏幕输出-->
    </outputs>
    <formats>
        <format id="main" format="%Date/%Time [%LEV] %Func:%Line %Msg%n"/>
    </formats>
</seelog>
`

func Init(logFile string) {
	if len(logFile) == 0{
		panic(errors.New("empty logFile"))
	}
	isExit := utils.FileExist(logFile)
	var logger log.LoggerInterface
	var err error
	if isExit{
		logger, err = log.LoggerFromConfigAsFile(logFile)
	}else {
		logger, err = log.LoggerFromConfigAsString(logFileString)
	}
	if err != nil {
		fmt.Println("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)
}
