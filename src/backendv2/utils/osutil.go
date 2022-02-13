package utils

import "os"

func GetEnvDefault(key, defaultVal string) string {
	if res, exist := os.LookupEnv(key); exist {
		return res
	}
	return defaultVal
}
