package decorator

import (
	"fmt"
	"time"
)

type Logger interface {
	Info(message string)
	Error(message string, err error)
}

type Log struct {
}

func (l Log) Info(message string) {
	fmt.Println("INFO", time.Now().Local().Format(time.RFC822), message)
}

func (l Log) Error(message string, err error) {
	fmt.Println("ERROR", time.Now().Local().Format(time.RFC822), message, err.Error())
}

func WithMessage(logger Logger) Logger {
	return &struct{ Logger }{
		Logger: logger,
	}
}

func WithError(logger Logger) Logger {
	return &struct{ Logger }{
		Logger: logger,
	}
}
