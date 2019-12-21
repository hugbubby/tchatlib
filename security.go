package tchatlib

import (
	"encoding/base64"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

type ServiceClaims struct {
	PublicKey ed25519.PublicKey `json:"public_key"`
	ServiceID string            `json:"service_id"`
	jwt.StandardClaims
}

type SigningMethodED25519 struct{}

func (s SigningMethodED25519) Verify(signingString, signature string, key interface{}) error {
	if edKey, ok := key.(ed25519.PublicKey); !ok {
	    return errors.New("invalid key; must be ed25519 public key")
	} else if len(edKey) != ed25519.PublicKeySize {
		return errors.New(fmt.Sprintf("invalid public key size; should be %d", ed25519.PublicKeySize))
	} else if signature_b, err := base64.StdEncoding.DecodeString(signature); err != nil {
		return errors.Wrap(err, "error decoding signature from base64")
	} else if !ed25519.Verify(edKey, []byte(signingString), signature_b) {
		return errors.New("invalid signature")
	} else {
        return nil
    }
}

func (s SigningMethodED25519) Sign(signingString string, key interface{}) (string, error) {
	var ret string
	var err error
	if edKey, ok := key.(ed25519.PrivateKey); !ok {
		err = errors.New("invalid key; must be ed25519 private key")
	} else if len(edKey) != ed25519.PrivateKeySize {
		err = errors.New(fmt.Sprintf("invalid private key size; should be %d", ed25519.PublicKeySize))
	} else {
		ret = base64.StdEncoding.EncodeToString(ed25519.Sign(edKey, []byte(signingString)))
	}
	return ret, err
}

func (s SigningMethodED25519) Alg() string {
	return "ED25519"
}
