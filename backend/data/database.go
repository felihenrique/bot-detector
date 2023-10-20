package data

import (
	"botdetector/app"
	"botdetector/utils"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mailru/go-clickhouse"
)

type database struct {
	conn *sql.DB
}

var Database = database{}

func (d *database) Connect(dsn string) {
	conn, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatalf("Can't connect to clickhouse: %s", err.Error())
	}

	d.conn = conn
}

func (d *database) Exec(query string, args ...any) (sql.Result, error) {
	return d.conn.Exec(query, args...)
}

func (d *database) InsertRequestLogs(logs []app.RequestLog) error {
	query := strings.Builder{}
	query.WriteString("INSERT INTO request_logs (player_id, ip, created_at, user_agent, is_bot) VALUES ")

	for index, entry := range logs {
		query.WriteString(fmt.Sprintf(
			"(%d, '%s', '%s', '%s', %d)",
			entry.PlayerId,
			entry.Ip.String(),
			utils.ToIsoString(entry.CreatedAt),
			entry.UserAgent,
			utils.BoolToInt(entry.IsBot),
		))

		if index < len(logs)-1 {
			query.WriteString(", ")
		}
	}
	finalQuery := query.String()
	_, err := d.conn.Exec(finalQuery)

	return err
}

func (d *database) Close() {
	d.conn.Close()
}
