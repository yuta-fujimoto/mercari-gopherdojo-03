package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: omikuji [port]")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	Run(flag.Args()[0])
}
