package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

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

	for _, p := range SortedList {
		fmt.Printf("IP:%s \t %s Port\n", p.IP, p.Port)

	}

}

func main() {

	gather()

}

func gather() {
	response, err := http.Get("https://pastebin.com/djMvDsTz")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// convert io page to string
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	newStr := buf.String()

	//setup reg match
	re := regexp.MustCompile("(\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\:\\d{3,5})")

	//find all that match
	match := re.FindAllString(newStr, -1)

	RAWProxyList = append(RAWProxyList, match...)

	SplitProxy()
}
