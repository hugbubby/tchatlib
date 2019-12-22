package tchatlib

import "golang.org/x/crypto/ed25519"

type Message struct {
	Sender  ed25519.PublicKey `json:"sender"`
	Content string            `json:"content"`
}
