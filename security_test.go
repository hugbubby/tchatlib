package tchatlib

import (
	"crypto/rand"
	"os"
	"testing"
	"testing/quick"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/ed25519"
)

var smname string

func TestMain(m *testing.M) {
	sm := SigningMethodED25519{}
	smname = sm.Alg()
	jwt.RegisterSigningMethod(sm.Alg(), func() jwt.SigningMethod { return sm })
	os.Exit(m.Run())
}

func TestPositiveVerify(t *testing.T) {
	sm := jwt.GetSigningMethod(smname)
	f := func(content string) bool {
		if pub, priv, err := ed25519.GenerateKey(rand.Reader); err != nil {
			t.Error(err)
		} else if sig_enc, err := sm.Sign(content, priv); err != nil {
			t.Error(err)
		} else if err := sm.Verify(content, sig_enc, pub); err != nil {
			t.Error(err)
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegativeVerify(t *testing.T) {
	sm := jwt.GetSigningMethod(smname)
	f := func(content string, fuzz byte, fuzz2 uint) bool {
		//Modified content
		pub, priv, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			t.Error(err)
		} else if sig_enc, err := sm.Sign(content, priv); err != nil {
			t.Error(err)
		} else {
			var newContent string
            if len(content) == 0 {
                newContent = string(fuzz)
            } else {
                index := fuzz2 % uint(len(content))
                if fuzz == content[index] {
                    fuzz++
                }
                newContent = content[:index] + string(fuzz) + content[index+1:]
            }
			return sm.Verify(newContent, sig_enc, pub) != nil
		}
		return false
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
