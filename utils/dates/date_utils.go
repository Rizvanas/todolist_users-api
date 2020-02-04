package dates

import "time"

const (
	apiDateLayout = "2006-01-10wT15:04:05Z"
)

// GetNow returns current UTC time in Time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetUTCTimeString returns current UTC time in string format
func GetUTCTimeString() string {
	return GetNow().Format(apiDateLayout)
}
