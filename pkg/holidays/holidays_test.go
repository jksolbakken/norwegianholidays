package holidays

import (
	"reflect"
	"testing"
	"time"
)

func TestYearsMustBeGreaterThanZero(t *testing.T) {
	invalidYears := map[int]time.Time{
		0:  {},
		-1: {},
	}
	for year := range invalidYears {
		_, err := EasterDay(year)
		if err == nil {
			t.Error("want an error, didn't get one")
		}
	}
}

func TestEasterDayForGivenYears(t *testing.T) {
	easterDays := map[int]time.Time{
		2032: NewDate(2032, time.March, 28).Truncate(time.Hour * 24),
		2023: NewDate(2023, time.April, 9).Truncate(time.Hour * 24),
		2022: NewDate(2022, time.April, 17).Truncate(time.Hour * 24),
		1985: NewDate(1985, time.April, 7).Truncate(time.Hour * 24),
		1950: NewDate(1950, time.April, 9).Truncate(time.Hour * 24),
		1854: NewDate(1854, time.April, 16).Truncate(time.Hour * 24),
	}

	for year := range easterDays {
		want := easterDays[year]
		got, _ := EasterDay(year)
		if want != got {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestAllPublicHolidaysForAGivenYear(t *testing.T) {
	calculated, _ := NorwegianHolidaysFor(2022)
	if !reflect.DeepEqual(calculated, norwegianHolidaysIn2022) {
		t.Errorf("got %v, wanted %v", calculated, norwegianHolidaysIn2022)
	}
}

var norwegianHolidaysIn2022 = []PublicHoliday{
	{NewDate(2022, time.January, 1), "1. nyttårsdag"},
	{NewDate(2022, time.April, 14), "Skjærtorsdag"},
	{NewDate(2022, time.April, 15), "Langfredag"},
	{NewDate(2022, time.April, 17), "1. påskedag"},
	{NewDate(2022, time.April, 18), "2. påskedag"},
	{NewDate(2022, time.May, 1), "Arbeidernes dag"},
	{NewDate(2022, time.May, 17), "Grunnlovsdagen"},
	{NewDate(2022, time.May, 26), "Kristi Himmelfartsdag"},
	{NewDate(2022, time.June, 5), "1. pinsedag"},
	{NewDate(2022, time.June, 6), "2. pinsedag"},
	{NewDate(2022, time.December, 25), "1. juledag"},
	{NewDate(2022, time.December, 26), "2. juledag"},
}
