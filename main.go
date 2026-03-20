package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"log"
	"net"
	"os"
)

var signer *Signer

func main() {
	loadKey()

	ln, err := net.Listen("tcp", ":26659")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Signer running on :26659")

	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}

func loadKey() {
	key := os.Getenv("SIGNER_PRIVATE_KEY")
	raw, _ := base64.StdEncoding.DecodeString(key)

	signer = NewSigner(ed25519.PrivateKey(raw))
	log.Println("Key loaded")
}

func handle(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	var req Request
	decoder.Decode(&req)

	log.Println("Received:", req.Type)

	switch req.Type {

	case "sign_vote":
		data, _ := base64.StdEncoding.DecodeString(req.SignBytes)

		sig := signer.Sign(data)

		encoder.Encode(Response{
			Status:    "ok",
			Signature: sig,
		})

	case "sign_vote_extension":
		data, _ := base64.StdEncoding.DecodeString(req.VoteExtBytes)

		sig := signer.SignVoteExtension(data)

		encoder.Encode(Response{
			Status:             "ok",
			ExtensionSignature: sig,
		})

	default:
		encoder.Encode(Response{
			Status: "error",
			Error:  "unknown type",
		})
	}
}
