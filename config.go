package tchatlib

import (
	"os"
)

type Config struct {
	Tor struct {
		ProxyAddress      string `json:"proxy_address"`
		ControllerAddress string `json:"controller_address"`
	} `json:"tor_config"`
	ServerAddress string `json:"server_address"`
	ReadCookie    string `json:"read_cookie"`
}

//Me being cheeky
var configDir = os.Getenv("HOME") + "/.config/tchatd"

func ConfigPath(filename string) string {
	return configDir + "/" + filename
}
