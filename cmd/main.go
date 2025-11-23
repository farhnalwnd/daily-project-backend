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

	err = db.AutoMigrate(
		&domain.ActivityMaster{},
		&domain.Activity{},
		&domain.ResourceCategory{},
		&domain.User{},
		&domain.Account{},
		&domain.Budgets{},
	)

	if err != nil {
		log.Fatalf("Gagal menjalankan Gorm AutoMigrate: %v", err)
	} else {
		fmt.Println("Berhasil menjalankan Gorm AutoMigrate")
	}
}
