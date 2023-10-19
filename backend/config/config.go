package config

import (
	"botdetector/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	DbUri   string
	Port    string
	IpsFile string
}

func readEnv() envConfig {
	godotenv.Load("./config/.env")

	cwd, _ := os.Getwd()
	ipsFile := "/res/ips.json.gz"

	if os.Getenv("IS_TESTING") == "1" {
		config := envConfig{
			IpsFile: cwd + "/.." + ipsFile,
		}

		return config
	}

	port := os.Getenv("PORT")
	if !utils.IsInt(port) {
		log.Fatal("Port needs to be a number")
	}

	dbUri := os.Getenv("CLICKHOUSE_URI")
	if len(dbUri) == 0 {
		log.Fatal("Clickhouse uri is required")
	}

	config := envConfig{
		DbUri:   os.Getenv("CLICKHOUSE_URI"),
		Port:    port,
		IpsFile: cwd + ipsFile,
	}

	return config
}

var Env = readEnv()
