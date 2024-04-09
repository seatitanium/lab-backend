package utils

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

/*
在 db 实例上执行语句，立即提交，并返回一个 sql.Result
*/
func DbExec(db *sqlx.DB, statement string, args ...any) (sql.Result, error) {
	tx := db.MustBegin()

	result, err := tx.Exec(statement, args)

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
func DbQuery(db *sqlx.DB, statement string, args ...any) (*sql.Rows, error) {
	tx := db.MustBegin()

	rows, err := tx.Query(statement, args)

	return rows, err
}

// 使用 tx.Get 尝试从数据库中提取一行单一数据，并将其写入到 *dest 中。如果找不到结果，将会返回一个 error
//
// 注意：参数 desc 应为指针
func DbGet(db *sqlx.DB, dest any, statement string, args ...any) error {
	tx := db.MustBegin()

	err := tx.Get(dest, statement, args)

	return err
}
