package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func AnyMatch(match string, strs ...string) bool {
	for _, str := range strs {
		if str == match {
			return true
		}
	}

	return false
}

func NoneMatch(match string, strs ...string) bool {
	return !AnyMatch(match, strs...)
}

func ParseTime(layout string, str string) (time.Time, error) {
	res, err := time.Parse(time.RFC3339, str)

	if err == nil {
		return res, nil
	}

	return time.Parse(layout, str)
}

func ParseTimeRFC3339(str string) (time.Time, error) {
	return ParseTime("2006-01-02T15:04:05Z", str)
}

func ParseTimeRFC3339Milli(str string) (time.Time, error) {
	return ParseTime("2006-01-02T15:04:05.999Z07:00", str)
}

func GetJsonFromData(dataName string, data interface{}) error {
	absPath, _ := filepath.Abs("data/" + dataName)
	bytes, readErr := os.ReadFile(absPath)

	if readErr != nil {
		return readErr
	}

	err := json.Unmarshal(bytes, data)

	if err != nil {
		return err
	}

	return nil
}
