package helper

import (
	"strings"
)

type ProxAddress struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

var sorted []ProxAddress

// GetUrls returns URL list to scrape
func GetUrls() ([]string, error) {

	var a []string
	a[0] = "test"
	a[1] = "ing"

	return a, nil
}

func SplitProxy(list []string) {

	for _, p := range list {
		//fmt.Printf("%s\n", p)
		s := strings.Split(p, ":")
		sorted = append(sorted, ProxAddress{s[0], s[1]})
	}

}
