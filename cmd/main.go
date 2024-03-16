package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"time"
	"weather-monitoring/models"
	"weather-monitoring/web"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	dsn := "root:Swadaya05@@tcp(localhost:3306)/weather_monitoring?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	defer func() {
		if sqlDB, err := DB.DB(); err != nil {
			panic(err)
		} else {
			sqlDB.Close()
		}
	}()

	DB.AutoMigrate(&models.Weather{})

	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for range ticker.C {
			water := rand.Intn(100) + 1
			wind := rand.Intn(100) + 1

			DB.Create(&models.Weather{Water: water, Wind: wind})

			updateJSON(water, wind)
		}
	}()

	r := web.SetupServer(DB)

	go func() {
		if err := r.Run(":8000"); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server is running on port 8000")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server...")
	ticker.Stop()
	fmt.Println("Server stopped")
}

func updateJSON(water, wind int) {
	data := map[string]map[string]int{
		"status": {
			"water": water,
			"wind":  wind,
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	err = ioutil.WriteFile("status.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("JSON file updated successfully")
}
