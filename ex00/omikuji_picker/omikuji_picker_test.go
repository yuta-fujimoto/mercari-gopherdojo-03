package omikuji_picker

import (
	"os"
	"testing"
	"time"
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

type TodayMock struct {
	year  int
	month time.Month
	day   int
}

func (t *TodayMock) Date() (int, time.Month, int) {
	return t.year, t.month, t.day
}

func TestPickInNewYear(t *testing.T) {
	ReadFile("../message.json")
	cases := []struct {
		date TodayMock
		want Fortune
		name string
	}{
		{
			TodayMock{year: 2022, month: time.January, day: 1},
			Daikichi,
			"1/1",
		},
		{
			TodayMock{year: 2022, month: time.January, day: 2},
			Daikichi,
			"1/2",
		},
		{
			TodayMock{year: 2022, month: time.January, day: 3},
			Daikichi,
			"1/3",
		},
	}
	for _, td := range cases {
		td := td
		t.Run(td.name, func(t *testing.T) {
			got := Pick(&td.date)
			if td.want != got.Fortune {
				t.Errorf("got: %d, want %d", got.Fortune, td.want)
			}
		})
	}
}

func TestPickNormal(t *testing.T) {
	t.Parallel()
	err := ReadFile("../message.json")
	if err != nil {
		t.Fatal("failed to read message file")
	}
	cases := []struct {
		date TodayMock
		name string
	}{
		{
			TodayMock{year: 2022, month: time.January, day: 4},
			"1/4",
		},
		{
			TodayMock{year: 2022, month: time.December, day: 31},
			"12/31",
		},
		{
			TodayMock{year: 2022, month: time.July, day: 3},
			"7/3",
		},
	}
	for _, td := range cases {
		td := td
		t.Run(td.name, func(t *testing.T) {
			t.Parallel()
			got := Pick(&td.date)
			validateFortune(t, got.Fortune.String())
		})
	}

}
func TestReadFile(t *testing.T) {
	err := ReadFile("nosuchfile")
	if err == nil {
		t.Error("expected error")
	}
	
	file, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal("failed to create tmp file")
	}
	defer os.Remove(file.Name())
	err = ReadFile(file.Name())
	if err == nil {
		t.Error("expected error")
	}
}
