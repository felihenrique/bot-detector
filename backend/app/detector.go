package app

import (
	"botdetector/config"
	"net"

	"github.com/zmap/go-iptree/iptree"
)

type detector struct {
	ips    *iptree.IPTree
	length int
}

var Detector = create()

func create() *detector {
	d := detector{}
	d.ips = iptree.New()
	d.loadIps(config.Env.IpsFile)

	return &d
}

func (d *detector) addIp(ipnet net.IPNet) {
	d.ips.Add(&ipnet, 0)
	d.length += 1
}

func (d *detector) IsBotAgent(agent string) (bool, error) {
	return false, nil
}

func (d *detector) IsBotIp(ip net.IP) (bool, error) {
	_, found, err := d.ips.Get(ip)

	if err != nil {
		return false, err
	}

	return found, nil
}
