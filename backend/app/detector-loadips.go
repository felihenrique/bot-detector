package app

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
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

	jsonData, err := io.ReadAll(gzipReader)
	if err != nil {
		log.Fatal(err)
	}

	var data []string
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
		return
	}

	for _, item := range data {
		_, ipnet, err := net.ParseCIDR(item)
		if err != nil {
			log.Fatal(err)
		}
		d.addIp(*ipnet)
	}
}
