package cfg

import (
	"os"

	"github.com/a2080016/lgpscan/internal/logger"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Lgp struct {
		Paths  []string `yaml:"paths"`
		Export struct {
			Clickhouse struct {
				Enabled bool   `yaml:"enabled"`
				Server  string `yaml:"server"`
				Port    int    `yaml:"port"`
			} `yaml:"clickhouse"`
			JSON struct {
				Enabled bool   `yaml:"enabled"`
				Path    string `yaml:"path"`
			} `yaml:"json"`
		} `yaml:"export"`
	} `yaml:"lgp"`
	Debug struct {
		PrintConfig    bool `yaml:"print config"`
		PrintLgfMaps   bool `yaml:"print lgf maps"`
		PrintLgpEvents bool `yaml:"print lgp events"`
	} `yaml:"debug"`
}

var AppConfig Config

func init() {

	cfgPath := `E:\go\lgpscan\config\config.yaml`

	file, err := os.Open(cfgPath)
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	defer file.Close()

	yDecoder := yaml.NewDecoder(file)
	yDecoder.Decode(&AppConfig)

	if AppConfig.Debug.PrintConfig {

		logger.InfLog.Println("Конфигурация:")
		logger.InfLog.Println("Пути к журналам:")
		for i := 0; i < len(AppConfig.Lgp.Paths); i++ {
			logger.InfLog.Println("    " + AppConfig.Lgp.Paths[i])
		}
		logger.InfLog.Println("Clickhouse:")
		logger.InfLog.Printf("    Enabled: %t", AppConfig.Lgp.Export.Clickhouse.Enabled)
		logger.InfLog.Printf("    Server: %v", AppConfig.Lgp.Export.Clickhouse.Server)
		logger.InfLog.Printf("    Port: %d", AppConfig.Lgp.Export.Clickhouse.Port)

		logger.InfLog.Println("JSON:")
		logger.InfLog.Printf("    Enabled: %t", AppConfig.Lgp.Export.JSON.Enabled)
		logger.InfLog.Printf("    Path: %v", AppConfig.Lgp.Export.JSON.Path)

		logger.InfLog.Println("Debug:")
		logger.InfLog.Printf("    PrintConfig: %t", AppConfig.Debug.PrintConfig)
		logger.InfLog.Printf("    PrintLgfMaps: %t", AppConfig.Debug.PrintLgfMaps)
		logger.InfLog.Printf("    PrintLgpEvents: %t", AppConfig.Debug.PrintLgpEvents)

		//fmt.Printf("%c", AppConfig.)
	}

}
