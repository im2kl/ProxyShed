package main

import (
	"fmt"

	"github.com/im2kl/ProxyShed/Client/admin"
	"github.com/im2kl/ProxyShed/Client/config"
	"github.com/im2kl/ProxyShed/Client/scraper"
)

//var RAWProxyList []string

func main() {

	config.Init()
	x := admin.GetURLList()

	s := scraper.Scrape(x)

	//y := proxydial.Test()

	//f := config.MaxThread()

	//time.Sleep(50 * time.Second)

	//for _, p := range s {
	fmt.Printf(s[0] + "\n")
	//}
}
