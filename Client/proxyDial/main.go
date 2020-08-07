package proxydial

// IPV4Address to identify socket type
type IPV4Address struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

// Test runs a test on the proxy to determine the type of proxy it is.
func Test() string {

	return "proxy type to return"
}

func dialHTTP() bool {

	return true
}

func dialSock4() bool {

	return true
}

func dialSock5() bool {

	return true
}
