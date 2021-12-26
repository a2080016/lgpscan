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
		JSONExport       bool   `yaml:"JSONExport"`
		JSONExportPath   string `yaml:"JSONExportPath"`
		YamlExport       bool   `yaml:"YamlExport"`
		YamlExportPath   string `yaml:"YamlExportPath"`
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

	//if AppConfig.Debug.PrintConfig {

	// logger.InfLog.Println("Конфигурация:")
	// logger.InfLog.Println("Пути к журналам:")
	// for i := 0; i < len(AppConfig.Lgp.Paths); i++ {
	// 	logger.InfLog.Println("    " + AppConfig.Lgp.Paths[i])
	// }
	// logger.InfLog.Println("Clickhouse:")
	// logger.InfLog.Printf("    Enabled: %t", AppConfig.Lgp.Export.Clickhouse.Enabled)
	// logger.InfLog.Printf("    Server: %v", AppConfig.Lgp.Export.Clickhouse.Server)
	// logger.InfLog.Printf("    Port: %d", AppConfig.Lgp.Export.Clickhouse.Port)

	// logger.InfLog.Println("JSON:")
	// logger.InfLog.Printf("    Enabled: %t", AppConfig.Lgp.Export.JSON.Enabled)
	// logger.InfLog.Printf("    Path: %v", AppConfig.Lgp.Export.JSON.Path)

	// logger.InfLog.Println("Debug:")
	// logger.InfLog.Printf("    PrintConfig: %t", AppConfig.Debug.PrintConfig)
	// logger.InfLog.Printf("    PrintLgfMaps: %t", AppConfig.Debug.PrintLgfMaps)
	// logger.InfLog.Printf("    PrintLgpEvents: %t", AppConfig.Debug.PrintLgpEvents)

	//fmt.Printf("%c", AppConfig.)
	//}

}
