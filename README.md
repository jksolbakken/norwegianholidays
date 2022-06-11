# norwegianholidays

Calculates all the Norewgian holidays for any given year.

Some of the holidays are at fixed dates, while others are relative to easter day. Easter day is [calculated](https://en.wikipedia.org/wiki/Date_of_Easter) using the Meeus/Jones/Butchers formula.

Usage:
```go
holidays, err := NorwegianHolidaysFor(2022)
```

Returns an array of `PublicHoliday`s rounded down to the day

```go
type PublicHoliday struct {
  When time.Time
  Name string
}
```

or an `error` if the supplied year is negative.
