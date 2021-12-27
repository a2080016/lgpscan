package status

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/a2080016/lgpscan/internal/logger"
	"gopkg.in/yaml.v3"
)

type StatusType struct {
	InfoBases map[string]struct {
		CurrentFile string `yaml:"CurrentFile"`
		CurrentPos  int64  `yaml:"CurrentPos"`
	} `yaml:"InfoBases,flow"`
}

var Status StatusType

func InitStatus(path string) {

	Status.InfoBases = make(map[string]struct {
		CurrentFile string "yaml:\"CurrentFile\""
		CurrentPos  int64  "yaml:\"CurrentPos\""
	})

	Status.InfoBases["www"] = struct {
		CurrentFile string "yaml:\"CurrentFile\""
		CurrentPos  int64  "yaml:\"CurrentPos\""
	}{
		CurrentFile: `222`,
		CurrentPos:  0,
	}

	data, err := yaml.Marshal(&Status)

	if err != nil {
		logger.ErrLog.Fatal(err)
	}

	err2 := ioutil.WriteFile(path, data, 0)

	if err2 != nil {

		logger.ErrLog.Fatal(err)
	}

	fmt.Println("data written")

}

func ReadStatus() {

	// Определяем текущий каталог для поиска файла статуса
	currentDirectory, err := os.Getwd()
	if err != nil {
		logger.ErrLog.Fatal(err)
	}

	statusPath := currentDirectory + `\data\status.yaml`
	logger.InfLog.Printf(statusPath)

	statusFile, err := os.Open(statusPath)
	if err != nil {
		InitStatus(statusPath)

	} else {
		yamlDecoder := yaml.NewDecoder(statusFile)
		yamlDecoder.Decode(&Status)
	}

}

func WriteStatus() {

	// Определяем текущий каталог для поиска файла статуса
	currentDirectory, err := os.Getwd()
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	logger.InfLog.Printf(currentDirectory)

	statusPath := currentDirectory + `\data\status.yaml`
	logger.InfLog.Printf(statusPath)

	statusFile, err := os.Open(statusPath)
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	defer statusFile.Close()

	yamlDecoder := yaml.NewDecoder(statusFile)
	yamlDecoder.Decode(&Status)

}
