package main

import (
	"github.com/im2kl/ProxyShed/Client/config"
)

type ProxAddress struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

var RAWProxyList []string

func main() {

	config.AdminDash.url = "proxysource.pnxbl.com"

}
