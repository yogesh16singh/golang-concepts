package main

import (
	"log"
	"os"
)

const (
	InfoLevel = iota
	WarningLevel
	ErrorLevel
)

type Logger struct {
	level       int
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

var logger *Logger

func init() {
	logger = &Logger{
		level:       InfoLevel,
		infoLogger:  log.New(os.Stdout, "INFO:", log.LstdFlags),
		warnLogger:  log.New(os.Stdout, "WARN:", log.LstdFlags),
		errorLogger: log.New(os.Stdout, "ERROR:", log.LstdFlags|log.Llongfile),
	}
}

func SetLevel(level int) {
	logger.level = level
}

func Info(message string) {
	if logger.level <= InfoLevel {
		logger.infoLogger.Println(message)
	}
}
func Warning(message string) {
	if logger.level <= WarningLevel {
		logger.warnLogger.Println(message)
	}
}
func Error(message string) {
	if logger.level <= ErrorLevel {
		logger.errorLogger.Println(message)
	}
}
func main() {

	Info("This is an info print")
	Warning("This is an Warning print")
	Error("This is an Error print")
	SetLevel(WarningLevel)
	Warning("This is an Warning print")

}
