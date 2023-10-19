package app

import (
	"net"
	"testing"
)

func TestLoadMmdb(t *testing.T) {
	ip1 := net.ParseIP("187.19.225.48")
	ip2 := net.ParseIP("50.16.251.200")

	result1, err1 := Detector.IsHostingIp(ip1)
	result2, err2 := Detector.IsHostingIp(ip2)

	if err1 != nil {
		t.Error(err1)
	}
	if err2 == nil {
		t.Error(err2)
	}

	t.Logf("IP1: %t", result1)
	t.Logf("IP2: %t", result2)
}
