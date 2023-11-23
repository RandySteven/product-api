package main

import (
	"log"
	"os"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
)

type Logger struct {
	UserLogger    *log.Logger
	ProductLogger *log.Logger
}

// PrintlnProduct implements interfaces.Logger.
func (l *Logger) PrintlnProduct(messages ...any) {
	l.ProductLogger.Println(messages...)
}

// PrintlnUser implements interfaces.Logger.
func (l *Logger) PrintlnUser(messages ...any) {
	l.UserLogger.Println(messages...)
}

var _ interfaces.Logger = &Logger{}

func (l *Logger) UserLog(file string) {
	logFile, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.SetFlags(log.LstdFlags)

	UserLogger := log.New(logFile, "user logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	l.UserLogger = UserLogger
}

func (l *Logger) ProductLog(file string) {
	logFile, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.SetFlags(log.LstdFlags)

	ProductLogget := log.New(logFile, "product logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	l.ProductLogger = ProductLogget
}
