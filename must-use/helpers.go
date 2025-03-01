package helpers

import (
	"fmt"
	"net/http"
	"os"
)

// Must is a helper function that panics if an error occurs
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// MustEnv reads an environment variable and panics if it’s missing
func MustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Environment variable %s is required", key))
	}
	return val
}

// MustQueryParam reads a query parameter from the URL and panics if it's missing
func MustQueryParam(r *http.Request, key string) string {
	val := r.URL.Query().Get(key)
	if val == "" {
		panic(fmt.Sprintf("Query parameter '%s' is required", key))
	}
	return val
}
