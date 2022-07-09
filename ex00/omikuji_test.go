package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
		t.Fatal("unexpected fortune value")
	}
}

func TestHandler(t *testing.T) {
	var v OmikujiJson

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	omikujiHandler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", rw.StatusCode)
	}
	body, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	if err := json.Unmarshal(body, &v); err != nil {
		t.Fatal(err.Error())
	}
	if v.Number < 1 || v.Number > 100 {
		t.Fatal("unexpected number")
	}
	validateFortune(t, v.Fortune)
}

// type TodayMock struct {
// 	year  int
// 	month time.Month
// 	day   int
// }

// func (t *TodayMock) Date() (int, time.Month, int) {
// 	return t.year, t.month, t.day
// }

// func TestGetFortunesIdx(t *testing.T) {
// 	cases := []struct {
// 		date TodayMock
// 		want int
// 		name string
// 	}{
// 		{
// 			TodayMock{year: 2022, month: time.January, day: 1},
// 			0,
// 			"1/1",
// 		},
// 		{
// 			TodayMock{year: 2022, month: time.January, day: 2},
// 			0,
// 			"1/2",
// 		},
// 		{
// 			TodayMock{year: 2022, month: time.January, day: 3},
// 			0,
// 			"1/3",
// 		},
// 	}
// 	for _, td := range cases {
// 		td := td
// 		t.Run(td.name, func(t *testing.T) {
// 			got := getFortuneIdx(&td.date)
// 			if td.want != got {
// 				t.Fatalf("got: %d, want %d", got, td.want)
// 			}
// 		})
// 	}
// }
