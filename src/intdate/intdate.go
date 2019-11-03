package intdate

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

// IntDate IntDate
type IntDate int

// NewIntDate NewIntDate
func NewIntDate(date int) (IntDate, error) {
	if !isValidDate(date) {
		return -1, errors.Errorf("Invalid date. date=%v", date)
	}

	return IntDate(date), nil
}

// ToPath ToPath
func (date IntDate) ToPath() string {
	return filepath.Join(
		fmt.Sprintf("%02d", toYear(int(date))),
		fmt.Sprintf("%02d", toMonth(int(date))),
		fmt.Sprintf("%02d", toDay(int(date))),
	)
}

// ToTime ToTime
func (date IntDate) ToTime() time.Time {
	return toTime(int(date))
}

func toYear(date int) int {
	return date / 10000
}

func toMonth(date int) int {
	y := toYear(date)
	if y == 0 {
		return -1
	}

	return (date / 100) % y
}

func toDay(date int) int {
	b := date / 100
	if b == 0 {
		return -1
	}

	return date % (date / 100)
}

func toTime(date int) time.Time {
	year := toYear(date)
	month := toMonth(date)
	day := toDay(date)

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func isValidDate(date int) bool {
	year := toYear(date)
	month := toMonth(date)
	day := toDay(date)
	t := toTime(date)

	return year == t.Year() && time.Month(month) == t.Month() && day == t.Day()
}
