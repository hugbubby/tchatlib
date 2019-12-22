package tchatlib

type DaemonConfig struct {
	PrivateServerAddress string `json:"private_server_address"`
	ClientCookie         string `json:"client_cookie"`
}

type TorConfig struct {
	ProxyAddress      string `json:"proxy_address"`
	ControllerAddress string `json:"controller_address"`
}
