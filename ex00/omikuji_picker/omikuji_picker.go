package omikuji_picker

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	fileName = "omikuji_picker/message.json"
)

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
		log.Fatal("fortune.String(): unknown value")
	}
	return ""
}


type omikuji struct {
	Number  int
	Fortune Fortune
	Msg     string
}

var omikujiList []omikuji
var newYearOmikujiList []omikuji

type Today interface {
	Date() (int, time.Month, int)
}

func init() {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	dec := json.NewDecoder(jsonFile)
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err.Error())
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
			log.Fatal(err.Error())
		}
		omikujiList = append(omikujiList,
			omikuji{Number: number, Fortune: Fortune(d.Fortune), Msg: d.Msg})
		if Fortune(d.Fortune) == Daikichi {
			newYearOmikujiList = append(newYearOmikujiList,
				omikuji{Number: number, Fortune: Fortune(d.Fortune), Msg: d.Msg})
		}
		number++
	}
}

func isNewYear(month time.Month, day int) bool {
	return month == time.January && (day >= 1 && day <= 3)
}

func Pick() omikuji {
	_, month, day := time.Now().Date()
	var picked omikuji
	if isNewYear(month, day) {
		picked = newYearOmikujiList[rand.Intn(len(newYearOmikujiList))]
	} else {
		picked = omikujiList[rand.Intn(len(newYearOmikujiList))]
	}

	return picked
}
