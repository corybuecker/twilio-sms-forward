package logging

import "log"

type Logger interface {
	Log(message string)
}

type StandardOutLogger struct{}

func (logger *StandardOutLogger) Log(message string) {
	log.Println(message)
}
