package app

type services struct{}

var Services = services{}

func (services) HydrateRequestData(data *RequestData) error {
	agentResult, err := Detector.IsBotAgent(data.UserAgent)
	ipRangeResult, err2 := Detector.IsBotIp(data.Ip)

	if err != nil {
		return err
	}
	if err2 != nil {
		return err
	}

	isBot := agentResult || ipRangeResult

	data.IsBot = isBot

	return nil
}
