package pkg

import "os"

func GetEnv(key, defaultValue string) string {
	// Get ENV By Key
	v := os.Getenv(key)

	// Check Exist Value or Not
	if v == "" {
		return defaultValue
	}

	// Return Value Env Not Null String
	return v
}
