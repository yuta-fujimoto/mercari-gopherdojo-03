package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"omikuji/omikuji_picker"
	"testing"
)

func validateFortune(t *testing.T, s string) {
	t.Helper()
	switch s {
	case "Dai-kichi":
		return
	case "Kichi":
		return
	case "Chuu-kichi":
		return
	case "Sho-kichi":
		return
	case "Sue-kichi":
		return
	case "Kyo":
		return
	case "Dai-kyo":
		return
	default:
		t.Error("unexpected fortune value")
	}
}

func TestHandler(t *testing.T) {
	t.Parallel()
	err := omikuji_picker.ReadFile("message.json")
	if err != nil {
		t.Fatal("failed to read message file")
	}
	var v OmikujiJson

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	omikujiHandler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", rw.StatusCode)
	}
	body, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Error("unexpected error")
	}
	if err := json.Unmarshal(body, &v); err != nil {
		t.Error(err.Error())
	}
	if v.Number < 1 || v.Number > 100 {
		t.Error("unexpected number")
	}
	validateFortune(t, v.Fortune)
}
