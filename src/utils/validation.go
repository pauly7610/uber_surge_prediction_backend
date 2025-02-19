package utils

import (
	"errors"
	"regexp"
)

func ValidateCoordinates(lat, lon float64) error {
	if lat < -90 || lat > 90 {
		return errors.New("invalid latitude")
	}
	if lon < -180 || lon > 180 {
		return errors.New("invalid longitude")
	}
	return nil
}

func ValidateTimeRange(start, end string) error {
	// Simple regex to check ISO 8601 format
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`)
	if !re.MatchString(start) || !re.MatchString(end) {
		return errors.New("invalid time format")
	}
	return nil
} 