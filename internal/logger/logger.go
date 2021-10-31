package logger

import (
	"log"
	"os"
)

var InfLog *log.Logger
var ErrLog *log.Logger

func init() {
	InfLog = log.New(os.Stdout, "INF\t", log.Ldate|log.Ltime)
	ErrLog = log.New(os.Stdout, "ERR\t", log.Ldate|log.Ltime)
}

func PrintInf(message string) {
	InfLog.Println(message)
}

func PrintErr(message string) {
	ErrLog.Println(message)
}
