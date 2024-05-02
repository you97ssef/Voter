package services

import (
	"fmt"
	"os"
	"time"
)

func NewFileLogger(src string) *Logger {
	return &Logger{
		shouldWriteToFile: true,
		src:               src,
	}
}

func NewDefaultLogger() *Logger {
	return &Logger{
		shouldWriteToFile: false,
		src:               "",
	}
}

type Logger struct {
	shouldWriteToFile bool
	src               string
}

func (l *Logger) writeToFile(message string) error {
	if l.shouldWriteToFile {
		file, error := os.OpenFile(l.src, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if error != nil {
			return error
		}
		
		defer file.Close()
		file.WriteString(message + "\n")
	}
	return nil
}

func (l *Logger) Info(message interface{}) error {
	message = fmt.Sprintf("%s [INFO] - %v", time.Now().Format("02/01/2006 15:04:05"), message)
	
	var err = l.writeToFile(fmt.Sprintf("%v", message))
	if err != nil {
		return err
	}

	fmt.Println("\033[32m" + fmt.Sprintf("%v", message) + "\033[0m")
	
	return nil
}

func (l *Logger) Alert(message interface{}) error {
	message = fmt.Sprintf("%s [ALERT] - %v", time.Now().Format("02/01/2006 15:04:05"), message)

	var err = l.writeToFile(fmt.Sprintf("%v", message))
	if err != nil {
		return err
	}

	fmt.Println("\033[31m" + fmt.Sprintf("%v", message) + "\033[0m")
	
	return nil
}

func (l *Logger) Warning(message interface{}) error {
	message = fmt.Sprintf("%s [WARNING] - %v", time.Now().Format("02/01/2006 15:04:05"), message)

	var err = l.writeToFile(fmt.Sprintf("%v", message))
	if err != nil {
		return err
	}
	
	fmt.Println("\033[33m" + fmt.Sprintf("%v", message) + "\033[0m")
	
	return nil
}
