package app

import (
	"botdetector/config"
	"net"
	"strings"

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

func (d *detector) IsBotIp(ip net.IP) (bool, error) {
	_, found, err := d.ips.Get(ip)

	if err != nil {
		return false, err
	}

	return found, nil
}

var crawlerAgents = []string{
	"okhttp",
	"googlebot",
	"adsbot-google",
	"headlesschrome",
	"facebookexternalhit",
	"python-requests",
	"aiohttp",
	"python-httpx",
	"Python-urllib",
	"python-urllib3",
	"nessus",
	"nessus::soap",
	"nessus/190402",
	"curl",
	"bingbot",
	"bingpreview",
	"semrushbot",
	"ahrefsbot",
	"chrome-lighthouse",
	"petalbot",
	"applebot",
	"pingdom.com_bot",
	"axios",
	"yandexbot",
	"ptst",
	"adsbot-google-mobile",
	"siteauditbot",
}

func (d *detector) IsBotAgent(agent string) bool {
	ag := strings.ToLower(agent)
	for _, crawler := range crawlerAgents {
		if strings.Contains(ag, crawler) {
			return true
		}
	}
	return false
}
