package models

// ProxySource has the URL and reg to use for scrapping
type ProxySource struct {
	Url string `json:"url"`
	Reg string `json:"reg"`
}

// ProxAddress of scrapped
type ProxAddress struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}
