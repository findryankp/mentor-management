package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() *AppConfig {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &AppConfig{
		DBUSERNAME: os.Getenv("DBUSERNAME"),
		DBPASS:     os.Getenv("DBPASS"),
		DBHOST:     os.Getenv("DBHOST"),
		DBPORT:     os.Getenv("DBPORT"),
		DBNAME:     os.Getenv("DBNAME"),
		JWTKEY:     os.Getenv("SECRETKEY"),
	}

}
