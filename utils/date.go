package utils

import (
	"strconv"
	"strings"
	"time"
)

// Date is a utility function that takes a date in string format(YYYY-MM-DD) and converts it into time.Time format
func Date(date string) (time.Time, error) {
	d := strings.Split(date, "-")
	t := time.Time{} // zeroth value of time is nil struct
	year, err := strconv.Atoi(d[0])
	if err != nil {
		return t, err
	}
	month, err := strconv.Atoi(d[1])
	if err != nil {
		return t, err
	}
	day, err := strconv.Atoi(d[2])
	if err != nil {
		return t, err
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}
