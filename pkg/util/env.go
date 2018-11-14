package util

import "os"

// GetEnvOrDefault reads an environment variable, or returns the given default of not found
func GetEnvOrDefault(envName, dfault string) string {

	env := os.Getenv(envName)
	if env == "" {
		return dfault
	}
	return env
}
