package datasource

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Exec interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type NamedQueryEntity interface {
	GetOrder() string
	GetLimit() string
	SetLimit()
	GetSize() int64
	GetOffset() int64
}

func NamedQueryListAndTotal[T any](t []*T, db *sqlx.DB, arg NamedQueryEntity, query, defaultSort, orderAlias string) (list []*T, total *int64) {

	countRow, err := db.NamedQuery("select count(*) "+query[strings.Index(query, "from "):], arg)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}

	listRows, err := db.NamedQuery(query, arg)
	if err != nil {
		panic(err)
	}
	for listRows.Next() {
		data := new(T)
		err = listRows.StructScan(data)
		if err != nil {
			panic(err)
		}
		t = append(t, data)
	}
	defer listRows.Close()

	list = make([]*T, 0, arg.GetSize())

	if *total > arg.GetOffset() {
		if arg.GetOrder() != "" {
			if orderAlias != "" {
				order := arg.GetOrder()
				query += order[:10] + orderAlias + "." + order[10:]
			} else {
				query += arg.GetOrder()
			}
		} else {
			query += " " + defaultSort
		}
		if arg.GetLimit() != "" {
			query += arg.GetLimit()
		}
		listRows, err := db.NamedQuery(query, arg)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			data := new(T)
			err := listRows.StructScan(data)
			if err != nil {
				panic(err)
			}
			list = append(list, data)
		}
		defer listRows.Close()
	}

	return
}
