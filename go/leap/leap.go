package leap

const testVersion = 2

// IsLeapYear determines whether a
// particular Gregorian calendar year is a leap year
func IsLeapYear(y int) bool {
	return y%4 == 0 && y%100 != 0 || y%400 == 0
}
