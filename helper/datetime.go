package helper

import (
	"time"
)

func ParseStringToTime(format, dateStr string, loc *time.Location) (*time.Time, error) {
	var (
		result time.Time
		err    error
	)

	if loc == nil {
		result, err = time.Parse(format, dateStr)
	} else {
		result, err = time.ParseInLocation(format, dateStr, loc)
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}
