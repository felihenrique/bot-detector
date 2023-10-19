package app

import (
	"net"
	"testing"
)

func TestLoadedIps(t *testing.T) {
	if Detector.ips == nil {
		t.Error("Ip tree should be initialized")
	}

	if Detector.length == 0 {
		t.Error("Ip file should be loaded")
	}
}

func TestDetectorIps(t *testing.T) {
	ip1 := net.ParseIP("187.19.225.48") // non-bot ip
	ip2 := net.ParseIP("50.16.251.200") // bot ip

	result1, err1 := Detector.IsBotIp(ip1)
	result2, err2 := Detector.IsBotIp(ip2)

	if result1 || err1 != nil {
		t.Error(err1)
	}
	if !result2 || err2 != nil {
		t.Error(err2)
	}

	t.Logf("IP1: %t", result1)
	t.Logf("IP2: %t", result2)
}

func TestDetectorAgent(t *testing.T) {
	botAgents := []string{
		"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/104.0.5112.101 Safari/537.36",
		"facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
		"Mozilla/5.0 (Linux; Android 7.0; Moto G (4)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4590.2 Mobile Safari/537.36 Chrome-Lighthouse",
	}

	for _, agent := range botAgents {
		isBot := Detector.IsBotAgent(agent)
		if !isBot {
			t.Error("Not detected a bot agent")
		}
	}
}
