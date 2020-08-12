package helper

import (
	"fmt"
	"strings"
	"sync"
)

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
		SortedList = append(SortedList, ProxAddress{s[0], s[1]})

	}


}
