package main

import (
	"encoding/json"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:26659")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	req := map[string]interface{}{
		"type": "sign_vote",
		"sign_bytes": "dGVzdA==",
		"height": 11,
		"round": 0,
		"step": "precommit",
	}

	json.NewEncoder(conn).Encode(req)

	var resp map[string]interface{}
	json.NewDecoder(conn).Decode(&resp)

	log.Println(resp)
}
