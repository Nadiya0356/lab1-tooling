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

	if err := viper.ReadInConfig(); err != nil {
		logger.Error().Err(err).Msg("Failed to read config")
		return
	}

	appName := viper.GetString("app_name")
	port := viper.GetInt("port")

	logger.Info().
		Str("app", appName).
		Int("port", port).
		Msg("Application started")

	// 🔹 логіка з лаби 1
	sum := internal.Add(5, 3)
	fmt.Println("Add:", sum)

	result, err := internal.Divide(10, 2)
	if err != nil {
		logger.Error().Err(err).Msg("Divide error")
		return
	}

	fmt.Println("Divide:", result)

	// обробка помилки
	if _, err := internal.Divide(5, 0); err != nil {
		logger.Warn().Err(err).Msg("Handled division by zero")
	}
}
