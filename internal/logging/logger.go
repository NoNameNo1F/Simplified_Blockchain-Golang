package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	_logFile *os.File
	_logger  *log.Logger
)

// init initializes the logger by opening a log file and setting up the logger.
func init() {
	var err error
	logFilePath := "data/logs.log"

	_logFile, err = os.OpenFile(
		logFilePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	SetLogOutput(true, true)
}

func SetLogOutput(printToStdout bool, writeToFile bool) {
	var writers []io.Writer

	if printToStdout {
		writers = append(writers, os.Stdout)
	}

	if writeToFile {
		writers = append(writers, _logFile)
	}

	if len(writers) == 0 {
		_logger = log.New(io.Discard, "", log.LstdFlags)
	} else {
		mw := io.MultiWriter(writers...)
		_logger = log.New(mw, "", log.LstdFlags)
	}
}

func Log(logType string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("[%s] - [%s]: %s", timestamp, strings.ToUpper(logType), message)

	_logger.Println(msg)
}

// Authors: https://github.com/NoNameNo1F/Simplified_Blockchain-Golang
