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

// Configuration - describes the configuration file
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
	LanguageFilter     struct {
		Enabled            bool     `json:"enabled"`
		ForbiddenLanguages []string `json:"forbidden_languages"`
		Message            string   `json:"message"`
		ErrorRateLimit     int64    `json:"error_rate_limit"`
	} `json:"language_filter"`
	DiscloseErrorInternals bool `json:"disclose_error_internals"`
}

// CurrentConfig - stores the current configuration
var CurrentConfig Configuration

// LoadConfig - loads configuration from a file and stores it in CurrentConfig
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

	// Check welcome message configuration
	if CurrentConfig.Welcome.Enabled && CurrentConfig.Welcome.Message == "" {
		panic("[!!!CONFIGURATION ERROR!!!] Welcome message is enabled, but not set.")
	}

	// Check language filter configuration
	langFilterConfig := CurrentConfig.LanguageFilter

	if langFilterConfig.Enabled {
		if langFilterConfig.Message == "" {
			panic("[!!!CONFIGURATION ERROR!!!] Language filter is enabled, but error message is not set.")
		}

		if len(langFilterConfig.ForbiddenLanguages) == 0 {
			panic("[!!!CONFIGURATION ERROR!!!] Language filter is enabled, but no languages are set.")
		}

		if langFilterConfig.ErrorRateLimit <= 0 {
			panic("[!!!CONFIGURATION ERROR!!!] " +
				"Language filter is enabled, but error rate limit whether is not set, or it is negative integer.")
		}
	}

	log.SetOutput(os.Stdout)
	log.Printf("Successfully loaded configuration from %s\n", filename)
}
