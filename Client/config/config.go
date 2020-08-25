package config

// appSettings is the api url for admin information
type appSettings struct {
	// URL admin url for dashboard
	url string
	// Token Connection token / license
	token string
	// maxT maximum number of threads to use
	maxT int
}

// Init initialize the config
func Init() {

	conf := appSettings{}

	//get from secure source or hardcode?
	conf.token = "asd"
	conf.url = "proxysource.pnxbl.com"
	conf.maxT = 10

}

func (t *appSettings) GetToken() string {
	return t.token
}

// MaxThread get max set threads
func (t *appSettings) MaxThread() int {
	return t.maxT
}
