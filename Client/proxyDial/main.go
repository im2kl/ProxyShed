package proxydial

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// IPV4Address to identify socket type
type IPV4Address struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

// Test runs a test on the proxy to determine the type of proxy it is.
func Test() string {

	return "proxy type to return"
}

func dialHTTP(x IPV4Address) bool {

	proxyURL, err := url.Parse("https://" + x.IP + ":" + x.Port)
	if err != nil {
		log.Fatal(err)
	}
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	myClient.Timeout = time.Second * 60
	request, err := http.NewRequest("GET", "https://www.google.com", nil)
	//request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")

	resp, err := myClient.Do(request)

	if err != nil {
		fmt.Printf("IP:%s \t %s Port\t FAILED\n", x.IP, x.Port)
		fmt.Printf(err.Error() + "\n")
		return false
	}
	defer resp.Body.Close()

	//httpcode := resp.Status

	//fmt.Printf("IP:%s \t %s Port\t %s\n", x.IP, x.Port, httpcode)
	return true
}

func dialSock4() bool {

	return true
}

func dialSock5() bool {

	return true
}
