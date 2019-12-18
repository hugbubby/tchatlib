package tchatlib

import (
	"crypto/rand"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ed25519"
)

//Me being cheeky
var configDir = os.Getenv("HOME") + "/.config/tchatd"

func ConfigPath(filename string) string {
	return configDir + "/" + filename
}

//Load private key from disk
func GetKeys() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	var pub ed25519.PublicKey
	var priv ed25519.PrivateKey
	b, err := ioutil.ReadFile(ConfigPath("id_ecc"))
	if err != nil {
		if os.IsNotExist(err) {
			pub, priv, err = ed25519.GenerateKey(rand.Reader)
			if err == nil {
				err = ioutil.WriteFile(ConfigPath("id_ecc"), priv, 0600)
				if err == nil {
					err = ioutil.WriteFile(ConfigPath("id_ecc.pub"), pub, 0644)
				}
			}
		}
	} else {
		priv = b
		b, err = ioutil.ReadFile(ConfigPath("id_ecc.pub"))
		if err == nil {
			pub = b
		}
	}
	return pub, priv, err
}
