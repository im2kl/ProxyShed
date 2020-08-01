package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

var SourceUrl = "proxysource.pnxbl.com"

type ProxySource struct {
	Url string `json:"url"`
	Reg string `json:"reg"`
}

type ProxAddress struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

var RAWProxyList []string

var SortedList []ProxAddress

func SplitProxy() {

	for _, p := range RAWProxyList {
		//fmt.Printf("%s\n", p)
		s := strings.Split(p, ":")
		SortedList = append(SortedList, ProxAddress{s[0], s[1]})

	}

	var wg sync.WaitGroup
	for _, p := range SortedList {
		fmt.Printf("IP:%s \t %s Port\n", p.IP, p.Port)

		wg.Add(1)

		go func() {
			proxyDial(p)
			wg.Done()
		}()

	}
	wg.Wait()
}

func proxyDial(x ProxAddress) {
	prox := x.IP + ":" + x.Port
	conn, err := net.Dial("tcp", prox)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
		return
	}
	fmt.Printf(conn.RemoteAddr().String() + "\n")

	defer conn.Close()
	fmt.Printf("No error###########################################################################################? \n")
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		req, err := rw.ReadString('\n')
		if err != nil {
			rw.WriteString("failed to read input")
			rw.Flush()
			return
		}

		rw.WriteString(fmt.Sprintf("Request received: %s", req))
		rw.Flush()
	}
}

func ProxyPing(x ProxAddress) {

	proxyUrl, err := url.Parse("https://" + x.IP + ":" + x.Port)
	if err != nil {
		log.Fatal(err)
	}
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	request, err := http.NewRequest("GET", "https://www.google.com", nil)
	//request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")

	resp, err := myClient.Do(request)

	if err != nil {
		fmt.Printf("IP:%s \t %s Port\t FAILED\n", x.IP, x.Port)
		fmt.Printf(err.Error() + "\n")
		return
	}
	defer resp.Body.Close()

	httpcode := resp.Status

	fmt.Printf("IP:%s \t %s Port\t %s\n", x.IP, x.Port, httpcode)
}

func main() {

	gather()

}

func gather() {
	response, err := http.Get("https://proxy-daily.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// convert io page to string
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	newStr := buf.String()

	//fmt.Printf(newStr)
	//setup reg match
	re := regexp.MustCompile("(\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\:\\d{3,5})")

	//find all that match
	match := re.FindAllString(newStr, -1)

	RAWProxyList = append(RAWProxyList, match...)

	SplitProxy()
}
