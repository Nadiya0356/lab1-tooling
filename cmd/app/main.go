package main

import (
	"fmt"
	"os"

	"github.com/Nadiya0356/lab1-tooling/internal"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func main() {
	// 🔹 логер
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// 🔹 конфіг
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error().Err(err).Msg("Config error")
		return
	}

	appName := viper.GetString("app_name")
	port := viper.GetInt("port")

	logger.Info().Msg("App started")

	// 🔹 твій код з лаби 1
	fmt.Println("App:", appName)
	fmt.Println("Port:", port)

	fmt.Println("Add:", internal.Add(5, 3))

	result, err := internal.Divide(10, 2)
	if err != nil {
		logger.Error().Err(err).Msg("Divide error")
	} else {
		fmt.Println("Divide:", result)
	}

	// оброблена помилка
	_, _ = internal.Divide(5, 0)
}
