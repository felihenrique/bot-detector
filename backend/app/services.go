package app

type services struct{}

var Services = services{}

func (services) CheckRequestData(data RequestData) bool {
	agentResult, _ := Detector.IsBotAgent(data.UserAgent)
	ipRangeResult, _ := Detector.IsHostingIp(data.Ip)

	return agentResult || ipRangeResult
}
