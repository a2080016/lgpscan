package main

import (
	"github.com/a2080016/lgpscan/internal/logger"
	"github.com/a2080016/lgpscan/pkg/lgpparser"
)

func init() {
	logger.PrintInf("LGP Scanner, начало работы")
}

func main() {
	lgpparser.ParseLgpTst()
	logger.PrintInf("LGP Scanner, завершение работы")
}
