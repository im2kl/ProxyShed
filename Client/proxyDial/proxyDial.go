package proxydial

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// IPV4Address to identify socket type
type IPV4Address struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

var SortedList []IPV4Address

// Test runs a test on the proxy to determine the type of proxy it is.
func Test() string {

	var wg sync.WaitGroup
	for _, p := range SortedList {
		fmt.Printf("IP:%s \t %s Port\n", p.IP, p.Port)

		wg.Add(1)

		go func() {
			dialHTTP(p)
			wg.Done()
		}()

		// going further add checks for socks4 and socsk 5 as an extra func. followed by IP lookup
		// socks4 : https://github.com/Bogdan-D/go-socks4
		// socks5 : https://play.golang.org/p/l0iLtkD1DV
	}
	wg.Wait()
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

/*
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

*/
