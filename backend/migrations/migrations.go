package main

import (
	"botdetector/config"
	"botdetector/data"
	"log"
)

func createRequestLogTable() {
	query := `
		CREATE TABLE IF NOT EXISTS request_logs (
				player_id Int32,
				ip String,
				created_at DateTime DEFAULT now(),
				user_agent String,
				is_bot UInt8
		) ENGINE = MergeTree()
		ORDER BY (created_at, player_id);
	`

	data.Database.Connect(config.Env.DbUri)
	_, err := data.Database.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %s", err.Error())
	}
	log.Default().Print("Table request_logs created")
}

func main() {
	createRequestLogTable()
}
