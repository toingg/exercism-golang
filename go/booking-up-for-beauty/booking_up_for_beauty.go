package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
// Schedule("7/25/2019 13:45:00")
// => 2019-07-25 13:45:00 +0000 UTC
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05" // "M/D/YYYY"
	// layout := "2/1/2006, 15:04:05" // "D/M/YYYY"

	// _ : mengabaikan output kedua yaitu output error, karena fungsi .Parse me-return 2 output sehingga jika t saja tidak bisa
	t, _ := time.Parse(layout, date)

	// alt handle error
	// t, err := time.Parse("1/02/2006 15:04:05", date)
	// if err != nil {
	// 	panic(err)
	// }
	// return t

	return t
}

// HasPassed returns whether a date has passed.
// HasPassed("July 25, 2019 13:45:00")
// => true
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

	t, _ := time.Parse(layout, date)

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
// IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
// => true
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"

	t, _ := time.Parse(layout, date)
	

	hour := t.Hour()

	return hour >= 12 && hour < 18

}

// Description returns a formatted string of the appointment time.
// Description("7/25/2019 13:45:00")
// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	layout := "1/2/2006 15:04:05"

	t, _ := time.Parse(layout, date)

	formattedTime := t.Format("Monday, January 2, 2006, at 15:04.")

	description := "You have an appointment on " + formattedTime

	return description
}

// AnniversaryDate returns a Time with this year's anniversary.
// AnniversaryDate()
// => 2020-09-15 00:00:00 +0000 UTC
func AnniversaryDate() time.Time {
	yearNow := time.Now().Year()
	annivDate := fmt.Sprintf("%d-09-15 00:00:00 +0000 UTC", yearNow)

	layout := "2006-01-02 15:04:05 -0700 MST"
	t, _ := time.Parse(layout, annivDate)
	
	return t
}
