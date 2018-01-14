package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	seconds := time.Duration(1e9) * time.Second
	t = t.Add(seconds)
	return t
}
