package main

import (
	"flag"
	"log"
	"math/rand"
	"omikuji/server"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Usage: omikuji [port]")
	}

	rand.Seed(time.Now().UnixNano())

	server.Run(flag.Args()[0])
}
