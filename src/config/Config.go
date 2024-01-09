/*
 * Config.go
 * Copyright (c) ti-bone 2023-2024
 */

package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Configuration struct {
	BotToken    string `json:"bot_token"`
	DbDSN       string `json:"db_dsn"`
	LogsID      int64  `json:"logs_id"`
	LogsTopicID int64  `json:"logs_topic_id"`
	Welcome     struct {
		Enabled bool   `json:"enabled"`
		Message string `json:"message"`
	} `json:"welcome"`
	IsProtectedDefault bool `json:"is_protected_default"`
}

var CurrentConfig Configuration

func LoadConfig(filename string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("error reading config: %v", err))
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(fmt.Sprintf("error closing config file: %v", err))
		}
	}(jsonFile)

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(fmt.Sprintf("error reading config: %v", err))
	}

	err = json.Unmarshal(byteValue, &CurrentConfig)
	if err != nil {
		panic(fmt.Sprintf("error unmarshalling config: %v", err))
	}

	if CurrentConfig.Welcome.Enabled && CurrentConfig.Welcome.Message == "" {
		panic("[!!!CONFIGURATION ERROR!!!] Welcome message is enabled, but not set.")
	}

	log.SetOutput(os.Stdout)
	log.Printf("Successfully loaded configuration from %s\n", filename)
}
