package mysql

import (
	"baize/app/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var mysqlDb *sqlx.DB

// Init 初始化MySQL连接
func Init() {
	var err error = nil
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", setting.Conf.MySQLConfig.User, setting.Conf.MySQLConfig.Password, setting.Conf.MySQLConfig.Host, setting.Conf.MySQLConfig.Port, setting.Conf.MySQLConfig.DB)
	mysqlDb, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	mysqlDb.SetMaxOpenConns(setting.Conf.MySQLConfig.MaxOpenConns)
	mysqlDb.SetMaxIdleConns(setting.Conf.MySQLConfig.MaxIdleConns)
	return
}

func GetMysqlDb() **sqlx.DB {
	return &mysqlDb
}

// Close 关闭MySQL连接
func Close() {
	_ = mysqlDb.Close()
}
