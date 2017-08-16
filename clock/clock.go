package clock

import (
	"fmt"
)

const testVersion = 4


// Create a struct to hold time info
type Clock struct {
	hour int
	minute int
}

// New : Create a new Clock object based on hour and minute
func New(hour int, minute int) Clock {
	var hrsAdj int
	var newMinute int
	var newHour int
	var e int

	newMinute, hrsAdj, e = NormalizeMins(minute)
	if e != 0 {
		panic("Error from NormalizeMins()")
	}

	newHour, e = NormalizeHrs(hour + hrsAdj)
	if e != 0 {
		fmt.Printf("Error from NormalizeHrs() with hour = %d\n", hour + hrsAdj)
		panic("NormalizeHrs()")
	}


	return Clock {
		hour: newHour,
		minute: newMinute,
	}
}

// NormalizeMins : normalizes minutes to x hours and y minutes < 60
func NormalizeMins(minutes int) (mins int, hoursAdj int, err int) {
	adj := 0
	var negHrMins int

	if 0 <= minutes && minutes < 60 {
		return minutes, 0, 0
	}

	if minutes >= 60 {
		return minutes % 60, minutes / 60, 0
	}

	if minutes < 0 {
		if minutes % 60 != 0 {
			negHrMins = 60 + (minutes % 60)
		} else {
			negHrMins = 0
		}
		if negHrMins > 0 {
			adj = -1 + adj
		}
		return negHrMins, (minutes / 60) + adj, 0
	}

	return 0, 0, -1
}

// NormalizeHrs : normalizes hours to a 24-hr schedule
func NormalizeHrs(hrs int) (hrsN int, err int) {
	if 0 <= hrs && hrs < 24 {
		return hrs, 0
	}

	if hrs >= 24 {
		return hrs % 24, 0
	}

	if hrs < 0 {
		if hrs % 24 == 0 {
			return 0, 0
		} else {
			return 24 + (hrs % 24), 0
		}
	}

	return 0, -1
}

// String : print a Clock object
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add : add minutes and return an updated Clock object
func (c Clock) Add(minutes int) Clock {
	currentMin := c.minute
	currentHr := c.hour

	return New(currentHr, currentMin + minutes)
}
