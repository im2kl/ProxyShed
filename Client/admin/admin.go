package admin

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	m "github.com/im2kl/ProxyShed/Client/models"
)

// rawlist to be returned after scrapping.
var rawlist []m.ProxySource

// GetURLList retreive latest url list for scraping
func GetURLList() []m.ProxySource {

	client := http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5000/api/v1/list", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-access-token", "dingding")

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		body, err := ioutil.ReadAll(resp.Body) //only intrested in body just now
		if err != nil {
			log.Fatalln(err)
		}

		//log.Println(string(body))

		json.Unmarshal(body, &rawlist)
	} else {
		log.Println(string(resp.StatusCode))
	}
	return rawlist
}
