package logger

import (
	"fmt"
	"log"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorOrange = "\033[33m"
)

type Logger struct {
	enableDebug bool
}

func NewLogger(enableDebug bool) *Logger {
	return &Logger{
		enableDebug: enableDebug,
	}
}

func (l *Logger) Info(message string, fields ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [INFO] %s", timestamp, message)

	if len(fields) > 0 {
		logMessage += fmt.Sprintf(" %v", fields)
	}

	log.Println(logMessage)
}

func (l *Logger) Error(message string, err error, fields ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [ERROR] %s", timestamp, message)

	if err != nil {
		logMessage += fmt.Sprintf(" - Error: %v", err)
	}

	if len(fields) > 0 {
		logMessage += fmt.Sprintf(" %v", fields)
	}

	log.Println(ColorRed + logMessage + ColorReset)
}

func (l *Logger) Warn(message string, fields ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [WARN] %s", timestamp, message)

	if len(fields) > 0 {
		logMessage += fmt.Sprintf(" %v", fields)
	}

	log.Println(ColorOrange + logMessage + ColorReset)
}

func (l *Logger) Debug(message string, fields ...interface{}) {
	if !l.enableDebug {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [DEBUG] %s", timestamp, message)

	if len(fields) > 0 {
		logMessage += fmt.Sprintf(" %v", fields)
	}

	log.Println(logMessage)
}

func (l *Logger) Success(message string, fields ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [SUCCESS] %s", timestamp, message)

	if len(fields) > 0 {
		logMessage += fmt.Sprintf(" %v", fields)
	}

	log.Println(ColorGreen + logMessage + ColorReset)
}
