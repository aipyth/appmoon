package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aipyth/appmoon/internal/database"
)

func main() {
	config := database.DBConfig{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     5432,
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	db, err := database.NewDB(config)
	if err != nil {
		fmt.Printf("unable to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	activity := database.Activity{
		ClassName: "SampleClassName",
		Title:     "SampleTitle",
		Timestamp: time.Now(),
	}

	err = db.WriteActivity(activity)
	if err != nil {
		fmt.Printf("unable to write activity to database: %v\n", err)
		return
	}

	fmt.Println("Activity written to database successfully")
}
