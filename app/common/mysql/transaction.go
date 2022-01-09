package mysql

import "database/sql"

type Transaction interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}
