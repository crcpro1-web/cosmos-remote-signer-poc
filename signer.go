package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
)

type Signer struct {
	privateKey ed25519.PrivateKey
}

func NewSigner(priv ed25519.PrivateKey) *Signer {
	return &Signer{privateKey: priv}
}

func (s *Signer) Sign(data []byte) string {
	sig := ed25519.Sign(s.privateKey, data)
	return base64.StdEncoding.EncodeToString(sig)
}

func (s *Signer) SignVoteExtension(data []byte) string {
	// TODO: important for CometBFT v0.38
	return s.Sign(data)
}

func ValidateSlashing(height, round int64, step string, lastHeight int64) error {
	if height < lastHeight {
		return errors.New("unsafe: lower height")
	}
	return nil
}
