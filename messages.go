package tchatlib

import (
	"time"

	"golang.org/x/crypto/ed25519"
)

type Message struct {
	Sender    ed25519.PublicKey `json:"sender"`
	Signature []byte            `json:"signature"`
	Content   string            `json:"content"`
}

type MessageLogFile struct {
	MessageLogs []struct {
		Message
		time.Time
	} `json:"message_logs"`
}
