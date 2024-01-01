/*
 * DataSource.go
 * Copyright (c) ti-bone 2023-2024
 */

package db

import (
	"feedbackBot/src/config"
	"feedbackBot/src/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Connection *gorm.DB

func Init() {
	var err error
	Connection, err = gorm.Open(postgres.Open(config.CurrentConfig.DbDSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %w", err))
	}

	log.SetOutput(os.Stdout)

	log.Println("Trying to auto-migrate users table...")
	err = Connection.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Errorf("failed to auto-migrate users table: %w", err))
	}

	log.Println("Auto-migrated users table, successfully connected to the DB.")
}
