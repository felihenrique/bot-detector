package data

import (
	"botdetector/domain"
	"botdetector/utils"
	"database/sql"
	"fmt"
	"log"
	"strconv"
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

func (d *database) InsertLogBatch(logs []domain.RequestLog) error {
	query := strings.Builder{}
	query.WriteString("INSERT INTO request_logs (player_id, ip, created_at, user_agent, is_bot) VALUES ")

	for index, entry := range logs {
		query.WriteString(fmt.Sprintf(
			"(%d, '%s', '%s', '%s', %d)",
			entry.PlayerId,
			entry.Ip.String(),
			entry.CreatedAt,
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

func (d *database) ReadBotsPerDay(params ReadParams) ([]BotsPerDayData, error) {
	var query = `
		SELECT COUNT(player_id) as BotQuantity,
		formatDateTime(created_at, '%Y-%m-%d') as Date,
		player_id as PlayerId
		FROM request_logs rl
		WHERE is_bot = 1 
		AND created_at BETWEEN ? AND ?
		{{PLAYER}}
		GROUP BY Date, PlayerId
		ORDER BY Date ASC
	`
	var queryParams = []any{
		params.StartDate,
		params.EndDate,
	}

	if params.PlayerId > 0 {
		query = strings.Replace(query, "{{PLAYER}}", "AND player_id = ? ", 1)
		queryParams = append(queryParams, strconv.Itoa(params.PlayerId))
	} else {
		query = strings.Replace(query, "{{PLAYER}}", " ", 1)
	}

	result, err := d.conn.Query(
		query,
		queryParams...,
	)

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var data []BotsPerDayData = []BotsPerDayData{}

	for result.Next() {
		var item BotsPerDayData

		if err := result.Scan(&item.BotQuantity, &item.Date, &item.PlayerId); err != nil {
			return nil, err
		}
		data = append(data, item)
	}

	return data, nil
}

func (d *database) Close() {
	d.conn.Close()
}
