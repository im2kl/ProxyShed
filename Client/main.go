package main

import (
	"fmt"

	"github.com/im2kl/ProxyShed/Client/config"
	proxydial "github.com/im2kl/ProxyShed/Client/proxyDial"
	"github.com/im2kl/ProxyShed/Client/scraper"
)

type ProxAddress struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

//var RAWProxyList []string

func main() {

	config.Init()

	x := scraper.Scrape()

	y := proxydial.Test()

	//f := config.MaxThread()

	//time.Sleep(50 * time.Second)

	for _, p := range x {
		fmt.Printf(p + "\n" + y + "\n")
	}
}
