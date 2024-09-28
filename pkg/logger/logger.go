package logger

import "sync"

type Common interface {
	Info(msg string)
	Infof(format string, args ...interface{})
	Fatal(err error, msg string)
	Fatalf(err error, format string, args ...interface{})
	Error(err error, msg string)
	Errorf(err error, format string, args ...interface{})
}

var instance Common
var once sync.Once

func Log() Common {
	if instance == nil {
		panic("logger not initialized")
	}

	return instance
}

func Setup(logger Common) {
	once.Do(func() {
		instance = logger
	})
}
