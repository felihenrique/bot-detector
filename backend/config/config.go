package config

import (
	"botdetector/utils"
	"os"

	"github.com/joho/godotenv"
)

type configStruct struct {
	ClickhouseUri string
	Port          string
}

func readEnv() configStruct {
	godotenv.Load("./config/.env")

	port := os.Getenv("PORT")
	if !utils.IsInt(port) {
		panic("Port needs to be a number")
	}

	dbUri := os.Getenv("CLICKHOUSE_URI")

	if len(dbUri) == 0 {
		panic("Clickhouse uri is required")
	}

	config := configStruct{
		ClickhouseUri: os.Getenv("CLICKHOUSE_URI"),
		Port:          port,
	}

	return config
}

var Config = readEnv()
