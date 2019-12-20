package tchatlib

import "golang.org/x/crypto/ed25519"

type Message struct {
	ServiceID string `json:"service_id"`
	Content   string `json:"content"`
}

type ContactList struct {
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	ServiceID string            `json:"service_id"`
	PubKey    ed25519.PublicKey `json:"pubkey"`
}
