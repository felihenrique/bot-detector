package data

type ReadParams struct {
	StartDate string
	EndDate   string
	PlayerId  int
}

type BotsPerDayData struct {
	BotQuantity int
	PlayerId    int
	Date        string
}
