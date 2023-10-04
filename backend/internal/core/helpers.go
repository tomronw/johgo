package core

func CheckRetries(retries int) bool {
	return retries >= 3
}
