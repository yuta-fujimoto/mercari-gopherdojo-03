package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func containsFortunes(arr [fortunesCnt]string, str string, t *testing.T) bool {
	t.Helper()
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func TestHandler(t *testing.T) {
	var v omikjiJson
	if err := omikujiPicker.init(); err != nil {
		t.Fatal(err.Error())
	}

	// httptest.NewRequest()
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
	if v.No < 1 || v.No > 100 {
		t.Fatal("unexpected number")
	}
	if !containsFortunes(fortunes, v.Fortune, t) {
		t.Fatal("unexpected fortunes")
	}
}

type TodayMock struct {
	year  int
	month time.Month
	day   int
}

func (t *TodayMock) Date() (int, time.Month, int) {
	return t.year, t.month, t.day
}

func TestGetFortunesIdx(t *testing.T) {
	cases := []struct {
		date TodayMock
		want int
		name string
	}{
		{
			TodayMock{year: 2022, month: time.January, day: 1},
			0,
			"1/1",
		},
		{
			TodayMock{year: 2022, month: time.January, day: 2},
			0,
			"1/2",
		},
		{
			TodayMock{year: 2022, month: time.January, day: 3},
			0,
			"1/3",
		},
	}
	for _, td := range cases {
		td := td
		t.Run(td.name, func(t *testing.T) {
			got := getFortuneIdx(&td.date)
			if td.want != got {
				t.Fatalf("got: %d, want %d", got, td.want)
			}
		})
	}
}
