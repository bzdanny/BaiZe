package dataUtil

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DB interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)

	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)

	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
}

type NamedQueryEntity interface {
	GetOrder() string
	GetLimit() string
	GetSize() int64
	GetOffset() int64
}

func NamedQueryListAndTotal[T any](db DB, t []*T, arg NamedQueryEntity, query, defaultSort, orderAlias string) (list []*T, total *int64) {

	countRow, err := db.NamedQuery("select count(*) "+query[strings.Index(query, " from "):], arg)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}

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
			err = listRows.StructScan(data)
			if err != nil {
				panic(err)
			}
			list = append(list, data)
		}
		defer listRows.Close()
	}

	return list, total
}

func NamedQueryList[T any](db DB, t []*T, arg NamedQueryEntity, query, defaultSort, orderAlias string) (list []*T) {
	list = make([]*T, 0, arg.GetSize())
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
		list = append(list, data)
	}
	defer listRows.Close()

	return list
}
