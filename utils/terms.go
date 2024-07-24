package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func GetTerms() []Term {
	data := make([]Term, 0)

	absPath, _ := filepath.Abs("data/terms.json")
	bytes, readErr := os.ReadFile(absPath)

	if readErr != nil {
		return []Term{}
	}

	err := json.Unmarshal(bytes, &data)

	if err != nil {
		return []Term{}
	}

	return data
}
