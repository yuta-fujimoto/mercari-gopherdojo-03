package main

import (
	"bufio"
	"math/rand"
	"os"
)

type omikuji struct {
	no  int
	msg string
}

type omikujiList struct {
	msgs [fortunesCnt][]omikuji
}

const (
	fileName = ".message"
)

func (p *omikujiList) init() error {
	fp, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var idx int
	var msg string
	no := 1
	for scanner.Scan() {
		idx = int(scanner.Text()[0]) - '0'
		msg = scanner.Text()[2:]
		p.msgs[idx] = append(p.msgs[idx], omikuji{no, msg})
		no++
	}

	return nil
}

func (p *omikujiList) pick(idx int) omikuji {
	max := len(p.msgs[idx])
	return p.msgs[idx][rand.Intn(max)]
}
