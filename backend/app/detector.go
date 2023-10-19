package app

import (
	"net"

	"github.com/zmap/go-iptree/iptree"
)

type detector struct {
	ips *iptree.IPTree
}

var Detector = new()

func new() *detector {
	d := detector{}
	d.ips = iptree.New()
	return &d
}

func (d *detector) addIp(ipnet net.IPNet) {
	d.ips.Add(&ipnet, 0)
}

func (d *detector) IsBotAgent(agent string) (bool, error) {
	return false, nil
}

func (d *detector) IsHostingIp(netip net.IP) (bool, error) {
	_, found, err := d.ips.Get(netip)

	if err != nil {
		return false, err
	}

	return found, nil
}
