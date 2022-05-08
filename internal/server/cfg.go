package server

import (
	"fmt"
	"os"

	"github.com/taalhach/fiber-web-boiler-plate/internal/server/internal/server/database"

	ini "github.com/nanitor/goini"
	"github.com/taalhach/fiber-web-boiler-plate/internal/server/internal/server/configs"
)

const (
	envKey = "SETTINGS"
)

var (
	MainConfigs *configs.MainConfig
)

func loadConfigs() (*configs.MainConfig, error) {

	file := os.Getenv(envKey)
	if file == "" {
		fmt.Printf("Missing env variable: %v", envKey)
		os.Exit(1)
	}

	dict, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	MainConfigs, err = configs.LoadMainConfig(dict)
	if err != nil {
		return nil, err
	}

	// must make database connection
	database.MustConnectDB(MainConfigs.Database)
	return MainConfigs, err
}
