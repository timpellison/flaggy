package main

import (
	"github.com/captechtimmy/flaggy/internal/persistence"
	"github.com/captechtimmy/flaggy/internal/server"
)

func main() {
	run()
}

func run() {
	store := persistence.NewInMemoryStore()
	srv := server.NewFlaggyService(store)
	e := srv.Start()
	if e != nil {
		panic(e)
	}
}
