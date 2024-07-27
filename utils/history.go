package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func GetHistoryTermPlayers() (map[string][]ServerPlayer, error) {

	var data map[string][]ServerPlayer

	// Note: should be relative to project root
	absPath, _ := filepath.Abs("data/history-term-players.json")
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

func GetHistoryLoginRecord(player string) *LoginRecord {

	targetRecord := LoginRecord{
		Id:         0,
		ActionType: true,
		CreatedAt:  time.Now(),
		Tag:        "",
		Player:     player,
	}

	containsFlag := false

	for _, t := range GetTerms() {
		if HistoryTermsContainsPlayer(t.Tag, player) {
			containsFlag = true
			parsedTime, err := ParseTimeRFC3339(t.StartAt + "T00:00:00Z")

			if err != nil {
				println(err.Error())
				continue
			}

			if parsedTime.Before(targetRecord.CreatedAt) {
				targetRecord.CreatedAt = parsedTime
				targetRecord.Tag = t.Tag
			}
		}
	}

	if !containsFlag {
		return nil
	}

	return &targetRecord
}
