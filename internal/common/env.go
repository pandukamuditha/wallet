package common

import (
	"errors"
	"os"
	"strconv"
)

var errEnvVar = errors.New("env: Variable not available for the provided key")

var JwtSigningSecret = []byte("password")

func GetEnvStr(key string) (string, error) {
	valString := os.Getenv(key)
	if valString == "" {
		return valString, errEnvVar
	}
	return valString, nil
}

func GetEnvInt(key string) (int, error) {
	valString, err := GetEnvStr(key)
	if err != nil {
		return -1, err
	}
	val, err := strconv.Atoi(valString)
	if err != nil {
		return -1, err
	}
	return val, nil
}
