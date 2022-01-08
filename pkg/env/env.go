package env

import (
	"fmt"
	"os"
	"strconv"
)

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}
	return value
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func MustMapEnv(target *string, key string) {
	value := MustGetEnv(key)
	*target = value
}

func GetPort(defaultPort int) int {
	if len(os.Getenv("PORT")) > 0 {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			fmt.Printf("cannot convert %q to number!\n", os.Getenv("PORT"))
		}
		return port
	}
	return defaultPort
}
