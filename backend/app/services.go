package app

import (
	"botdetector/data"
	"botdetector/domain"
	"botdetector/utils"
	"time"
)

type services struct{}

var Services = services{}

func (services) hydrateRequestLog(item *domain.RequestLog) error {
	agentResult := Detector.IsBotAgent(item.UserAgent)
	ipRangeResult, err := Detector.IsBotIp(item.Ip)

	if err != nil {
		return err
	}

	if item.CreatedAt == "" {
		item.CreatedAt = utils.ToIsoString(time.Now())
	}

	isBot := agentResult || ipRangeResult
	item.IsBot = isBot

	return nil
}

func (services) ReadLogs(params ReadLogsParams) ([]data.BotsPerDayData, error) {
	return data.Database.ReadBotsPerDay(data.ReadParams(params))
}

func (services) WriteLog(item *domain.RequestLog) error {
	err := Services.hydrateRequestLog(item)
	if err != nil {
		return err
	}

	var itemCpy = *item
	data.AsyncWriter.Add(itemCpy)

	return nil
}
