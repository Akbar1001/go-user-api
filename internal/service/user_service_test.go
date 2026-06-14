package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {

	dob := time.Date(
		1990,
		time.May,
		10,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	age := CalculateAge(dob)

	if age <= 0 {
		t.Errorf(
			"invalid age returned: %d",
			age,
		)
	}
}