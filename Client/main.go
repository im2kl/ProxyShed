package main

import (
	"fmt"

	"github.com/im2kl/ProxyShed/Client/config"
	"github.com/im2kl/ProxyShed/Client/scraper"
)

type ProxAddress struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

//var RAWProxyList []string

func main() {

	//config.AdminDash.url = "proxysource.pnxbl.com"

	config.Init()

	//config.AdminDash.URL = "asda"

	x := scraper.Scrape()

	//time.Sleep(50 * time.Second)

	for _, p := range x {
		fmt.Printf(p + "\n")
	}
}
