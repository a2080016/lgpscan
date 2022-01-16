package status

import (
	"os"

	"github.com/a2080016/lgpscan/internal/logger"
	"gopkg.in/yaml.v3"
)

type StatusType struct {
	InfoBases map[string]struct {
		CurrentFile string `yaml:"CurrentFile"`
		CurrentPos  int64  `yaml:"CurrentPos"`
	} `yaml:"InfoBases"`
}

var CurrStatus StatusType

// Инициализирует файл статусов
//
func InitStatus(path string) {

	CurrStatus.InfoBases = make(map[string]struct {
		CurrentFile string "yaml:\"CurrentFile\""
		CurrentPos  int64  "yaml:\"CurrentPos\""
	})

	data, err := yaml.Marshal(&CurrStatus)

	if err != nil {
		logger.ErrLog.Fatal(err)
	}

	err2 := os.WriteFile(path, data, 0666)

	if err2 != nil {
		logger.ErrLog.Fatal(err)
	}

	logger.InfLog.Printf("InitStatus data written")
}

func ReadStatus() {

	// Определяем текущий каталог для поиска файла статуса
	currentDirectory, err := os.Getwd()
	if err != nil {
		logger.ErrLog.Fatal(err)
	}

	statusPath := currentDirectory + `\data\status.yaml`

	statusFile, err := os.Open(statusPath)
	if err != nil {
		InitStatus(statusPath)
	} else {
		yamlDecoder := yaml.NewDecoder(statusFile)
		yamlDecoder.Decode(&CurrStatus)
	}
	statusFile.Close()

}

func WriteStatus() {

	// Определяем текущий каталог для поиска файла статуса
	currentDirectory, err := os.Getwd()
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	logger.InfLog.Printf(currentDirectory)

	statusPath := currentDirectory + `\data\status.yaml`

	data, err := yaml.Marshal(&CurrStatus)

	if err != nil {
		logger.ErrLog.Fatal(err)
	}

	err2 := os.WriteFile(statusPath, data, 0666)

	if err2 != nil {
		logger.ErrLog.Fatal(err2)
	}

	logger.InfLog.Printf("WriteStatus data written")

}
