package main

import (
	"flag"
	"log"
	"math/rand"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Usage: omikuji [port]")
	}

	rand.Seed(time.Now().UnixNano())

	Run(flag.Args()[0])
}
