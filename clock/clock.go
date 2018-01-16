package clock

import (
	"strconv"
)

type Clock struct {
	timestamp int
	hour      int
	minute    int
}

func New(hour, minute int) Clock {
	if hour >= 24 {
		hour = hour % 24
	}
	timestamp := hour*60 + minute
	for timestamp < 0 {
		timestamp += 60 * 24
	}
	if timestamp >= 60*24 {
		timestamp %= 60 * 24
	}
	hour = timestamp / 60
	minute = timestamp - hour*60
	return Clock{timestamp: hour*60 + minute, hour: hour, minute: minute}
}

func format(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

func (c Clock) String() string {
	return format(c.hour) + ":" + format(c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
