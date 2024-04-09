package utils

import "github.com/jmoiron/sqlx"

func GetUserByUsername(db *sqlx.DB, username string) (*DbUser, error) {
	tx := db.MustBegin()

	var result DbUser

	if err := tx.Get(&result, "SELECT * FROM `seati_users` WHERE `username`=?", username); err != nil {
		return nil, err
	}

	return &result, nil
}
