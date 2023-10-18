package app

import (
	"net"
)

type detector struct{}

func (d detector) isBotAgent(agent string) bool {
	return false
}

func (d detector) isDatacenterIp(ip net.IP) bool {
	return false
}

var Detector = detector{}
