package tchatlib

import (
	"crypto/rand"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ed25519"
)


//Load private key from disk
func GetKeys(keypath string) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	var pub ed25519.PublicKey
	var priv ed25519.PrivateKey
	b, err := ioutil.ReadFile(keypath)
	if err != nil {
		if os.IsNotExist(err) {
			pub, priv, err = ed25519.GenerateKey(rand.Reader)
			if err == nil {
				err = ioutil.WriteFile(keypath, priv, 0600)
				if err == nil {
					err = ioutil.WriteFile(keypath, pub, 0644)
				}
			}
		}
	} else {
		priv = b
		b, err = ioutil.ReadFile(keypath)
		if err == nil {
			pub = b
		}
	}
	return pub, priv, err
}
