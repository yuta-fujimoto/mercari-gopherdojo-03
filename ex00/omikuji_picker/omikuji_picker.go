package omikuji_picker

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Mock
type Today interface {
	Date() (int, time.Month, int)
}

// Fortune
type Fortune int

const (
	Daikichi Fortune = iota
	Kichi
	Chuukichi
	Shokichi
	Suekichi
	Kyo
	Daikyo
)

func (f Fortune) String() string {
	switch f {
	case Daikichi:
		return "Dai-kichi"
	case Kichi:
		return "Kichi"
	case Chuukichi:
		return "Chuu-kichi"
	case Shokichi:
		return "Sho-kichi"
	case Suekichi:
		return "Sue-kichi"
	case Kyo:
		return "Kyo"
	case Daikyo:
		return "Dai-kyo"
	default:
		fmt.Fprintln(os.Stderr, "fortune.String(): unknown value")
		os.Exit(1)
	}
	return ""
}

// omikuji
type omikuji struct {
	Number  int
	Fortune Fortune
	Msg     string
}

var omikujiList []omikuji
var newYearOmikujiList []omikuji

func ReadFile(fn string) error {
	jsonFile, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	dec := json.NewDecoder(jsonFile)
	_, err = dec.Token()
	if err != nil {
		return err
	}

	type Data struct {
		Fortune int
		Msg     string
	}
	var d Data
	number := 1
	for dec.More() {
		err := dec.Decode(&d)
		if err != nil {
			return err
		}
		omikujiList = append(omikujiList,
			omikuji{Number: number, Fortune: Fortune(d.Fortune), Msg: d.Msg})
		if Fortune(d.Fortune) == Daikichi {
			newYearOmikujiList = append(newYearOmikujiList,
				omikuji{Number: number, Fortune: Fortune(d.Fortune), Msg: d.Msg})
		}
		number++
	}
	return nil
}

func isNewYear(month time.Month, day int) bool {
	return month == time.January && (day >= 1 && day <= 3)
}

func Pick(today Today) omikuji {
	_, month, day := today.Date()
	var picked omikuji

	if isNewYear(month, day) {
		picked = newYearOmikujiList[rand.Intn(len(newYearOmikujiList))]
	} else {
		picked = omikujiList[rand.Intn(len(omikujiList))]
	}
	return picked
}
