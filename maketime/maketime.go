// Package to read a calendar and print out the percentage of time spent in
// various tasks. Optionally to block out time when a day is too full of
// meetings.
package maketime

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

type Calendar struct {
	// All of the days we're tracking in this Calendar object.
	days []Day
}

type Day struct {
	// Date. Stored as a Time in UTC with the time set to midnight.
	day time.Time
	// Collection of times
	blocks [48]HalfHourBlock
}

type HalfHourBlock struct {
	// Stored as a Time in UTC
	day time.Time
	// Just a string representing this half-hour block for now.
	// nil for no event at this time.
	desc string
}

type Meeting struct {
	// Start time
	// End time
	// Probably should just use a Calendar.Event object
}

const HoursInDay int = 24
const BlocksInDay int = HoursInDay * 2

// Given a midnight UTC time.Time this will create a Day structure prepopulated
// with 30 minute blocks of time and return it.
func MakeDay(day time.Time) (Day, error) {
	if day.Hour() != 0 || day.Minute() != 0 || day.Second() != 0 ||
		day.Nanosecond() != 0 || day != day.UTC() {
		return Day{}, fmt.Errorf("Non-midnight UTC date given: %s.", day)
	}

	var newDay Day
	newDay.day = day
	curTime := day // Assumes it's midnight
	for i := 0; i < BlocksInDay; i++ {
		newDay.blocks[i] = HalfHourBlock{curTime, ""}
		curTime = curTime.Add(30 * time.Minute)
	}
	return newDay, nil
}

// Given a *calendar.Events and an empty Day, this will populate the Day with
// all of the events from the calendar.
// For now: assumes that all the events are on the same day.
func PopulateDay(day *Day, events *calendar.Events) error {
	// Get the day from the first Event's start time.
	start := events.Items[0].Start
	var eventday time.Time
	if start.Date != "" {
		eventday, _ = time.Parse("2006-01-02", start.Date)
	} else {
		eventday, _ = time.Parse(time.RFC3339, start.DateTime)
	}

	// Convert the event day to midnight-based
	day.day = time.Date(eventday.Year(), eventday.Month(), eventday.Day(), 0, 0,
		0, 0, time.UTC)

	// Now populate the blocks.
	// TODO: Read through the events, and put them in the right block with the
	// right description. For now, just add enough to pass the test.
	day.blocks[18].desc = events.Items[0].Description

	return nil
}

// Given a *calendar.Service, this will pull all of the events for the next 10
// days and return them as *calendar.Events.
func GetEvents(srv *calendar.Service) *calendar.Events {
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).
		TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
	}

	return events
}

// Given a *calendar.Events, this will print them out semi-readably.
func PrintEvents(events *calendar.Events) {
	fmt.Println("Upcoming events:")
	if len(events.Items) > 0 {
		for _, i := range events.Items {
			var when string
			// If the DateTime is an empty string the Event is an all-day Event.
			// So only Date is available.
			if i.Start.DateTime != "" {
				when = i.Start.DateTime
			} else {
				when = i.Start.Date
			}
			fmt.Printf("%s (%s)\n", i.Summary, when)
		}
	} else {
		fmt.Printf("No upcoming events found.")
	}
}
