package tchatlib

type Config struct {
	Tor struct {
		ProxyAddress      string `json:"proxy_address"`
		ControllerAddress string `json:"controller_address"`
	} `json:"tor_config"`
	PublicServerAddress  string `json:"public_server_address"`
	PrivateServerAddress string `json:"private_server_address"`
	ReadCookie           string `json:"read_cookie"`
}
