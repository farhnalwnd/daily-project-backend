package main

import (
	"dailyTracking/config"
	"dailyTracking/internal/core/domain"
	"dailyTracking/internal/infrastructure/database"
	"fmt"
	"log"
)

func main() {
	// Program sederhana untuk menguji Air
	fmt.Println("=======================================")
	fmt.Println("Aplikasi Go BERJALAN via Air!")
	fmt.Println("=======================================")

	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	log.Println("mulai migrasi database")

	err = db.AutoMigrate(
		&domain.ActivityMaster{},
		&domain.Activity{},
		&domain.ResourceCategory{},
		&domain.User{},
	)

	if err != nil {
		log.Fatalf("Gagal menjalankan Gorm AutoMigrate: %v", err)
	}
	log.Println("Migrasi skema Gorm selesai!")
}
