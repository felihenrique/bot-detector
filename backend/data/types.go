package data

type ReadParams struct {
	StartDate string
	EndDate   string
	PlayerId  int
}

type BotsPerDayData struct {
	BotQuantity int    `json:"bot_quantity"`
	PlayerId    int    `json:"player_id"`
	Date        string `json:"date"`
}
