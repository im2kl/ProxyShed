package config

// adminDash is the api url for admin information
type adminDash struct {
	// URL admin url for dashboard
	url string
	// Token Connection token / license
	token string
}

// Init initialize the config
func Init() {

	conf := adminDash{}

	//get from secure source or hardcode?
	conf.token = "asd"
	conf.url = "proxysource.pnxbl.com"

}
