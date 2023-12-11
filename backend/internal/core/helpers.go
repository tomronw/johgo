package core

// CheckRetries check if over retries
func CheckRetries(retries int) bool {
	return retries >= 3
}

// ValidateString validate string pointers
func ValidateString(str *string, defaultReturn string) string {
	if str != nil {
		return *str
	}
	return defaultReturn
}

// ValidateFloat64 validate float64 pointers
func ValidateFloat64(fl64 *float64, defaultReturn float64) float64 {
	if fl64 != nil {
		return *fl64
	}
	return defaultReturn
}
