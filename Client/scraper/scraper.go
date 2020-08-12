package scraper

import (
	"bytes"
	"log"
	"net/http"
	"regexp"
)

// ProxySource has the URL and reg to use for scrapping
type ProxySource struct {
	Url string `json:"url"`
	Reg string `json:"reg"`
}

// RAWProxyList to be returned after scrapping.
var RAWProxyList []string

// Scrape returns a list of scrapped proxies as a ip:port format
func Scrape() {
	response, err := http.Get("https://gist.github.com/futurex189/3769289")
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

}
