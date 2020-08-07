package config

// AdminURL is the api url for admin information
type AdminURL struct {
	url string
}

// LicenseToken is the auth token for the API/ also used as the license activation token
type LicenseToken struct {
	token string
}
