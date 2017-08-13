package leap

const testVersion = 3

// IsLeapYear : Returns true/false depending on whether or not year is leap year
func IsLeapYear(year int) bool {
	if (year % 4 == 0) {
		if (year % 100 != 0) {
			return true
		} else {
			if (year % 400 == 0) {
				return true
			}
			return false
		}
	}

	return false
}
