package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"omikuji/omikuji_picker"
)


type OmikujiJson struct {
	Number  int    `json:"number"`
	Fortune string `json:"fortune"`
	Message string `json:"message"`
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	picked := omikuji_picker.Pick()
	result := OmikujiJson{Number: picked.Number, Fortune: picked.Fortune.String(), Message: picked.Msg}

	json, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-TIPS", "It's a joke, of course!!!")
	fmt.Fprint(w, string(json))
}

func Run(port string) {
	http.HandleFunc("/", omikujiHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err.Error())
	}
}
