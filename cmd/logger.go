package main

import "log"

type Logger interface {
	Log(string)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Log(s string) {
	log.Println(s)
}
