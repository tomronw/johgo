package core

func CheckRetries(retries int) bool {

	if retries >= 3 {
		return false
	}
	return true
}
