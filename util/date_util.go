package util

import "time"

func IsFutureDate(date time.Time) bool {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	payrollYear, payrollMonth, _ := date.Date()
	return payrollYear > currentYear || (payrollYear == currentYear && payrollMonth > currentMonth)
}
