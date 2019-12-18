package tchatlib

type Config struct {
	Tor struct {
		ProxyAddress      string `json:"proxy_address"`
		ControllerAddress string `json:"controller_address"`
	} `json:"tor_config"`
	ServerAddress string `json:"server_address"`
	ReadCookie    string `json:"read_cookie"`
}
