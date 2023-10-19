package app

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"log"
	"net"
	"os"

	"github.com/samber/lo"
)

func (d *detector) loadIps(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	scanner := bufio.NewScanner(gzipReader)
	for scanner.Scan() {
		line := scanner.Text()

		var data []string
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			continue
		}
		lo.ForEach(data, func(item string, index int) {
			_, ipnet, err := net.ParseCIDR(item)
			if err != nil {
				log.Fatal(err)
			}
			d.addIp(*ipnet)
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
