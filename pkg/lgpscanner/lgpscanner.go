package lgpscanner

import (
	"github.com/a2080016/lgpscan/pkg/lgpparser"
)

func Scan() {

	infoBase := "ds_estate"
	lgpparser.ScanLgpFiles(infoBase)

}
