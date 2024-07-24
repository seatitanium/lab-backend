package static

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type HistoryTermPlayers struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

func GetHistoryTermPlayers() (map[string][]HistoryTermPlayers, error) {

	var data map[string][]HistoryTermPlayers

	// Note: should be relative to project root
	absPath, _ := filepath.Abs("static/history-term-players.json")
	bytes, readErr := os.ReadFile(absPath)

	if readErr != nil {
		return nil, readErr
	}

	err := json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func HistoryTermsContainsPlayer(tag string, player string) bool {
	history, err := GetHistoryTermPlayers()

	if err != nil {
		println(err.Error())
		return false
	}

	for _, x := range history[tag] {
		if x.Name == player {
			return true
		}
	}

	return false
}
