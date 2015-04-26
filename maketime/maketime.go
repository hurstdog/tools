// Package to read a calendar and print out the percentage of time spent in
// various tasks. Optionally to block out time when a day is too full of
// meetings.
package maketime

import (
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
	"time"
)

type Day struct {
	// Date. Stored as a Time in UTC with the time set to midnight.
	day time.Time
	// Collection of times
	blocks [48]HalfHourBlock
}

type HalfHourBlock struct {
	// Stored as a Time in UTC
	day time.Time
	// Pointer to event taking up that time
	// If nil, then this block is free
}

type Meeting struct {
	// Start time
	// End time
	// Probably should just use a Calendar.Event object
}

const HoursInDay int = 24
const BlocksInDay int = HoursInDay * 2

func MakeDay(day time.Time) (Day, error) {
	if day.Hour() != 0 || day.Minute() != 0 || day.Second() != 0 ||
		day.Nanosecond() != 0 || day != day.UTC() {
		return Day{}, fmt.Errorf("Non-midnight UTC date given: %s.", day)
	}

	var newDay Day
	newDay.day = day
	curTime := day // Assumes it's midnight
	for i := 0; i < BlocksInDay; i++ {
		newDay.blocks[i] = HalfHourBlock{curTime}
		curTime = curTime.Add(30 * time.Minute)
	}
	return newDay, nil
}

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
