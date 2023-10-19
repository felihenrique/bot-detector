package app

type services struct{}

var Services = services{}

func (services) HydrateRequestData(data *RequestData) error {
	agentResult := Detector.IsBotAgent(data.UserAgent)
	ipRangeResult, err := Detector.IsBotIp(data.Ip)

	if err != nil {
		return err
	}

	isBot := agentResult || ipRangeResult
	data.IsBot = isBot

	return nil
}
