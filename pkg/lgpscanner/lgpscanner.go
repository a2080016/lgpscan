package lgpscanner

import "github.com/a2080016/lgpscan/pkg/lgpparser"

func Scan() {

	lgpparser.ScanLgp()

	//for i := 0; i < len(cfg.AppConfig); i++ {

	// file, err := os.Open(cfg.AppConfig.Lgp.Paths[i])
	// if err != nil {
	// 	logger.ErrLog.Println("Нет такого файла")
	// } else {
	// 	logger.InfLog.Println(file.Name())

	// 	fstat, err2 := file.Stat()
	// 	if err2 != nil {
	// 		logger.ErrLog.Println("ошибка stat")
	// 	}

	// 	logger.InfLog.Println(fstat.IsDir())
	// 	logger.InfLog.Println(fstat.ModTime())
	// 	logger.InfLog.Println(fstat.Mode())
	// 	logger.InfLog.Println(fstat.Name())
	// 	logger.InfLog.Println(fstat.Size())
	// 	logger.InfLog.Println(fstat.Sys())
	//}

}