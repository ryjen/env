package context

import (
	"os"
	"strings"
)

var (
	// the environment variable name to determine context
	identifier = "ENV_CONTEXT"
)

// Get returns the current context
func Get() string {
	val, ok := os.LookupEnv(identifier)
	if ok {
		return val
	}
	return identifier
}

// Is compares the value as the context insensitive to case
func Is(value string) bool {
	return strings.EqualFold(os.Getenv(identifier), value)
}

// Set configures the name of the env variable to determine context with
// Setting an empty string effectively turns of contextual environments
func Set(value string) {
	identifier = value
}

// IsEnv returns true if the identifier exists as an environment variable
func IsEnv() bool {
	_, ok := os.LookupEnv(identifier)
	return ok
}
