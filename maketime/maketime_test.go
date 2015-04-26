package maketime

import (
	"fmt"
	"testing"
	"time"

	"google.golang.org/api/calendar/v3"
)

func TestMakeDay(t *testing.T) {
	expected := getTestDay()

	result, err := MakeDay(expected.day)
	if err != nil {
		t.Errorf("Got error %s from MakeDay()", err)
	}

	if result != expected {
		printDayComparison(&result, &expected)
		t.Error("MakeDay failed.")
	}
}

func TestPopulateDay(t *testing.T) {
	e := NewEvent("First Test", "2009-11-10 09:00:00", "2009-11-10 09:30:00")
	events := calendar.Events{}
	events.Items = []*calendar.Event{&e}
	result, err := MakeDay(time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Got error %s from MakeDay", err)
	}

	err = PopulateDay(&result, &events)
	if err != nil {
		t.Errorf("Got error %s from PopulateDay", err)
	}

	expected := getTestDay()
	// Set up a meeting for 9-9:30am
	expected.blocks[18].desc = "First Test"
	if result != expected {
		printDayComparison(&result, &expected)
		t.Error("MakeDay failed.")
	}
}

// Given a datetime in "YYYY-MM-DD HH:MM:SS" format, returns an rfc3339
// formatted string.
// Fails on error.
func rfc3339(datetime string) string {
	// Mon Jan 2 15:04:05 MST 2006
	format := "2006-01-02 15:04:05"
	t, err := time.Parse(format, datetime)
	if err != nil {
		fmt.Printf("ERROR: Couldn't parse time %s with format %s", datetime, format)
	}
	return t.Format(time.RFC3339)
}

// Creates an event with the given description, start, and end times. The start
// and end times need to be in "YYYY-MM-DD HH:MM:SS" format.
func NewEvent(desc string, start string, end string) calendar.Event {
	e := calendar.Event{}
	e.Start = &calendar.EventDateTime{DateTime: rfc3339(start)}
	e.End = &calendar.EventDateTime{DateTime: rfc3339(end)}
	e.Description = desc

	return e
}

// Pretty-prints a comparison between the two days, just saying the values that
// differ.
func printDayComparison(day1 *Day, day2 *Day) {
	if *day1 == *day2 {
		fmt.Print("day1 and day2 are equal.\n")
		return
	}

	if day1.day != day2.day {
		fmt.Printf("Day mismatch: %s != %s\n", day1.day, day2.day)
		return
	}

	// Looks like we have to print out the block differences...
	for i, day1value := range day1.blocks {
		if day1value != day2.blocks[i] {
			fmt.Printf("Day block %d: [%s]@%s != [%s]@%s\n", i, day1value.desc, day1value.day,
				day2.blocks[i].desc, day2.blocks[i].day)
		}
	}
}

// Returns the canonical test day, of 10 november 2009, fully populated with
// half hour blocks.
func getTestDay() Day {
	testDay := time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)

	// Ugly, but it makes it really clear how it's supposed to work.
	blocks := [48]HalfHourBlock{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 0, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 1, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 1, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 2, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 2, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 3, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 3, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 4, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 4, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 5, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 5, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 6, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 6, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 7, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 7, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 8, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 8, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 9, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 9, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 10, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 10, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 11, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 11, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 12, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 12, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 13, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 13, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 14, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 14, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 15, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 15, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 16, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 16, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 17, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 17, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 18, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 18, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 19, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 19, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 20, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 20, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 21, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 21, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 22, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 22, 30, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), ""},
		{time.Date(2009, time.November, 10, 23, 30, 0, 0, time.UTC), ""},
	}

	return Day{testDay, blocks}
}
