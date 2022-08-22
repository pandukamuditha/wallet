package common

import (
	"errors"
	"os"
	"strconv"
)

var envVarErr = errors.New("env: Variable not available for the provided key")

func GetEnvStr(key string) (string, error) {
	valString := os.Getenv(key)
	if valString == "" {
		return valString, envVarErr
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