package utils

import "database/sql"

/*
在 db 实例上执行语句，并返回一个 sql.Result
*/
func DbExec(db *sql.DB, statement string, parameters string) (sql.Result, error) {
	stmt, err := db.Prepare(statement)
	defer MustPanic(stmt.Close())
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(parameters)
	return result, err
}

/*
在 db 实例上进行语句查询，并返回一个 *sql.Rows
*/
func DbQuery(db *sql.DB, statement string, parameters string) (*sql.Rows, error) {
	stmt, err := db.Prepare(statement)
	defer MustPanic(stmt.Close())
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(parameters)
	return rows, err
}
