package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type omikjiJson struct {
	No      int    `json:"number"`
	Fortune string `json:"fortune"`
	Message string `json:"message"`
}

const (
	fortunesCnt = 7
)

type Today interface {
	Date() (int, time.Month, int)
}

var (
	fortunes      = [fortunesCnt]string{"Dai-kichi", "Kichi", "Chuu-kichi", "Sho-kichi", "Sue-kichi", "Kyo", "Dai-kyo"}
	omikujiPicker = omikujiList{}
)

func getFortuneIdx(t Today) int {
	_, month, day := t.Date()
	if month == time.January && (day >= 1 && day <= 3) {
		return 0
	}
	return rand.Intn(len(fortunes))
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	idx := getFortuneIdx(time.Now())
	picked := omikujiPicker.pick(idx)
	result := omikjiJson{No: picked.no, Fortune: fortunes[idx], Message: picked.msg}

	json, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// https://developer.mozilla.org/ja/docs/Web/HTTP/Headers: Custom header
	w.Header().Set("X-TIPS", "It's a joke, of course!!!")
	fmt.Fprint(w, string(json))
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: omikuji [port]")
		os.Exit(1)
	}

	if err := omikujiPicker.init(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", omikujiHandler)
	if err := http.ListenAndServe(":"+flag.Arg(0), nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
