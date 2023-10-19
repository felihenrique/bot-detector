package config

import (
	"botdetector/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	DbUri       string
	Port        string
	IpsFilePath string
}

func readEnv() envConfig {
	godotenv.Load("./config/.env")

	port := os.Getenv("PORT")
	if !utils.IsInt(port) {
		log.Fatal("Port needs to be a number")
	}

	dbUri := os.Getenv("CLICKHOUSE_URI")
	if len(dbUri) == 0 {
		log.Fatal("Clickhouse uri is required")
	}

	config := envConfig{
		DbUri:       os.Getenv("CLICKHOUSE_URI"),
		Port:        port,
		IpsFilePath: "./res/ips.csv.gz",
	}

	return config
}

var Env = readEnv()
