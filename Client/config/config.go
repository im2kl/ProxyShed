package config

// appSettings is the api url for admin information
type appSettings struct {
	// URL admin url for dashboard
	Url string
	// Token Connection token / license
	Token string
	// maxT maximum number of threads to use
	MaxT int
}

// Init initialize the config
func Init() {

	conf := appSettings{}

	//get from secure source or hardcode?
	conf.Token = "asd"
	conf.Url = "http://127.0.0.1:5000"
	conf.MaxT = 10

}

func GetToken() string {
	t := appSettings{}

	return t.Token
}

// MaxThread get max set threads
func MaxThread() int {
	t := appSettings{}
	return t.MaxT
}

// MaxThread get max set threads
func GetURL() string {
	t := appSettings{}
	return t.Url
}
