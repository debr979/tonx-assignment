package utils

import (
	"io"
	"log"
	"os"
)

type logger struct {
}

var Logger logger

func (r *logger) LogOutput(errStr ...interface{}) {
	fileName := `log.out`
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	ErrLog := log.New(io.MultiWriter(logFile, os.Stderr), `[ERROR]`, log.Ldate|log.Ltime|log.Llongfile)

	ErrLog.Println(errStr...)
}
