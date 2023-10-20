package app

import "time"

type services struct{}

var Services = services{}

func (services) HydrateRequestLog(data *RequestLog) error {
	agentResult := Detector.IsBotAgent(data.UserAgent)
	ipRangeResult, err := Detector.IsBotIp(data.Ip)

	if err != nil {
		return err
	}

	if (data.CreatedAt == time.Time{}) {
		data.CreatedAt = time.Now()
	}

	isBot := agentResult || ipRangeResult
	data.IsBot = isBot

	return nil
}
