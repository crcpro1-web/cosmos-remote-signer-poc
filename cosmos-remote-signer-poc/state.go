package main

import "sync"

type State struct {
	LastHeight int64
}

var state = &State{}
var mu sync.Mutex
