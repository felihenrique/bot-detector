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
