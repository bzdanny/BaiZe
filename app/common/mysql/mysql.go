package mysql

import (
	"baize/app/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var masterMysqlDb *sqlx.DB

// Init 初始化MySQL连接
func Init() {
	var err error = nil
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", setting.Conf.MySQLConfig.User, setting.Conf.MySQLConfig.Password, setting.Conf.MySQLConfig.Host, setting.Conf.MySQLConfig.Port, setting.Conf.MySQLConfig.DB)
	masterMysqlDb, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	masterMysqlDb.SetMaxOpenConns(setting.Conf.MySQLConfig.MaxOpenConns)
	masterMysqlDb.SetMaxIdleConns(setting.Conf.MySQLConfig.MaxIdleConns)
	return
}

func GetMasterMysqlDb() *sqlx.DB {
	return masterMysqlDb
}

// Close 关闭MySQL连接
func Close() {
	_ = masterMysqlDb.Close()
}
