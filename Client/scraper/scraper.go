package scraper

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"

	m "github.com/im2kl/ProxyShed/Client/models"
)

// rawlist to be returned after scrapping.
var rawlist []string

// Scrape returns a list of scrapped proxies as a ip:port format
func Scrape(proxlist []m.ProxySource) []string {

	for _, p := range proxlist { // turn to go routine
		fmt.Println(p.Reg + "\t" + p.Url)

		response, err := http.Get(p.Url)
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
		re := regexp.MustCompile(p.Reg)

		//find all that match
		match := re.FindAllString(newStr, -1)

		rawlist = append(rawlist, match...)

		//if rawlist != nil {
		//	return rawlist
		//}

	}

	return rawlist
}
