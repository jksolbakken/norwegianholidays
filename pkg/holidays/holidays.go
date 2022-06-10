package holidays

import (
	"fmt"
	"math"
	"time"
)

// EasterDay Meeus/Jones/Butchers formula
func EasterDay(year int) (time.Time, error) {
	if year <= 0 {
		return time.Time{}, fmt.Errorf("cannot do negative years: %d", year)
	}
	a := year % 19
	b := math.Floor(float64(year / 100.0))
	c := year % 100
	d := math.Floor(b / 4)
	e := int(b) % 4
	f := math.Floor((b + 8) / 25)
	g := math.Floor((b - f + 1) / 3)
	h := int(19*a+int(b)-int(d)-int(g)+15) % 30
	i := math.Floor(float64(c / 4.0))
	k := c % 4
	l := (32 + 2*int(e) + 2*int(i) - h - k) % 7
	m := int(math.Floor(float64((a + 11*h + 22*l) / 451)))
	month := math.Floor(float64((h + l - 7*m + 114) / 31))
	dayOfMonth := ((h + l - 7*int(m) + 114) % 31) + 1
	return NewDate(year, time.Month(month), dayOfMonth), nil
}

type PublicHoliday struct {
	When time.Time
	Name string
}

func NewDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 1, 1, 1, 1, time.UTC).Truncate(time.Hour * 24)
}

func NorwegianHolidaysFor(year int) ([]PublicHoliday, error) {
	easterDayDate, err := EasterDay(year)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	easterDayHoliday := PublicHoliday{easterDayDate, "1. påskedag"}
	return []PublicHoliday{
		{NewDate(year, time.January, 1), "1. nyttårsdag"},
		{easterDayHoliday.plusDays(-3), "Skjærtorsdag"},
		{easterDayHoliday.plusDays(-2), "Langfredag"},
		easterDayHoliday,
		{easterDayHoliday.plusDays(1), "2. påskedag"},
		{NewDate(year, time.May, 1), "Arbeidernes dag"},
		{NewDate(year, time.May, 17), "Grunnlovsdagen"},
		{easterDayHoliday.plusDays(39), "Kristi Himmelfartsdag"},
		{easterDayHoliday.plusDays(49), "1. pinsedag"},
		{easterDayHoliday.plusDays(50), "2. pinsedag"},
		{NewDate(year, time.December, 25), "1. juledag"},
		{NewDate(year, time.December, 26), "2. juledag"},
	}, nil
}

func (h *PublicHoliday) plusDays(n int) time.Time {
	return h.When.Add(time.Duration(n*24) * time.Hour).Truncate(time.Hour * 24)
}
