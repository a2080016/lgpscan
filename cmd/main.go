package main

import (
	"github.com/a2080016/lgpscan/internal/logger"
	"github.com/a2080016/lgpscan/pkg/lgpscanner"
)

func init() {
	logger.PrintInf("LGP Scanner, начало работы")
}

func main() {

	lgpscanner.Scan()
	logger.PrintInf("LGP Scanner, завершение работы")
}
