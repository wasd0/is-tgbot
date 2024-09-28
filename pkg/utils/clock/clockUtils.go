package clock

import "time"

func GetUtc(timePtr *time.Time) *time.Time {
	if timePtr == nil {
		return nil
	}

	utc := timePtr.UTC()
	return &utc
}
