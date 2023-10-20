package app

type services struct{}

var Services = services{}

func (services) HydrateRequestLog(data *RequestLog) error {
	agentResult := Detector.IsBotAgent(data.UserAgent)
	ipRangeResult, err := Detector.IsBotIp(data.Ip)

	if err != nil {
		return err
	}

	isBot := agentResult || ipRangeResult
	data.IsBot = isBot

	return nil
}
