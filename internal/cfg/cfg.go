package cfg

import (
	"os"

	"github.com/a2080016/lgpscan/internal/logger"
	"gopkg.in/yaml.v3"
)

type ConfigType struct {
	Debug     bool `yaml:"Debug"`
	InfoBases map[string]struct {
		LgpPath          string `yaml:"LgpPath"`
		ClickHouseExport bool   `yaml:"ClickHouseExport"`
		ClickHouseServer string `yaml:"ClickHouseServer"`
		ClickHousePort   int    `yaml:"ClickHousePort"`
		JsonExport       bool   `yaml:"JsonExport"`
		JsonExportPath   string `yaml:"JsonExportPath"`
		YamlExport       bool   `yaml:"YamlExport"`
		YamlExportPath   string `yaml:"YamlExportPath"`
		CsvExport        bool   `yaml:"CsvExport"`
		CsvExportPath    string `yaml:"CsvExportPath"`
		Show             bool   `yaml:"Show"`
	} `yaml:"InfoBases,flow"`
}

var Config ConfigType

func init() {

	// Определяем текущий каталог для поиска конфиг. файла
	currentDirectory, err := os.Getwd()
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	logger.InfLog.Printf(currentDirectory)

	configPath := currentDirectory + `\config\config.yaml`
	logger.InfLog.Printf(configPath)

	configFile, err := os.Open(configPath)
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	defer configFile.Close()

	yamlDecoder := yaml.NewDecoder(configFile)
	yamlDecoder.Decode(&Config)

	logger.InfLog.Println(Config)

}
