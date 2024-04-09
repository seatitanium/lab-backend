package utils

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

/*
在 db 实例上执行语句，并返回一个 sql.Result
*/
func DbExec(db *sqlx.DB, statement string, parameters ...any) (sql.Result, error) {
	tx := db.MustBegin()

	result, err := tx.Exec(statement, parameters)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return result, err
}

/*
在 db 实例上进行语句查询，并返回一个 *sql.Rows
*/
func DbQuery(db *sqlx.DB, statement string, parameters ...any) (*sql.Rows, error) {
	tx := db.MustBegin()

	rows, err := tx.Query(statement, parameters)

	return rows, err
}
