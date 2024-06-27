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

func init() {
	var err error

	_logFile, err = os.OpenFile(
		"logs.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	mw := io.MultiWriter(os.Stdout, _logFile)
	_logger = log.New(mw, "", 0)
}

func Log(logType string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("[%s] - [%s]: %s", timestamp, strings.ToUpper(logType), message)

	_logger.Println(msg)
}
