package utils

import (
	"time"
)

func GetHistoryTermPlayers() (map[string][]ServerPlayer, error) {

	var data map[string][]ServerPlayer

	err := GetJsonFromData("history-term-players.json", &data)

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
