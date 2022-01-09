package datasource

import (
	"baize/app/setting"
	"fmt"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var masterDb *sqlx.DB
var slaveDb []*sqlx.DB

// Init 初始化MySQL连接
func Init() {
	datasource := setting.Conf.Datasource
	master := datasource.Master
	var err error = nil
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", master.User, master.Password, master.Host, master.Port, master.DB)
	masterDb, err = sqlx.Connect(master.DriverName, dsn)
	if err != nil {
		panic(err)
	}
	masterDb.SetMaxOpenConns(master.MaxOpenConns)
	masterDb.SetMaxIdleConns(master.MaxIdleConns)

	slave := datasource.Slave
	count := slave.Count
	if count > 0 {
		slaveDb = make([]*sqlx.DB, count)
		for i := 0; i < count; i++ {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", slave.Users[i], slave.Passwords[i], slave.Hosts[i], slave.Ports[i], slave.DBs[i])
			slaveDb[i], err = sqlx.Connect(slave.DriverName, dsn)
			if err != nil {
				panic(err)
			}
			slaveDb[i].SetMaxOpenConns(master.MaxOpenConns)
			slaveDb[i].SetMaxIdleConns(master.MaxIdleConns)
		}
	}
	return
}

// GetMasterDb 获取主数据源
func GetMasterDb() *sqlx.DB {
	return masterDb
}

// GetSlaveDb 获取从数据源
func GetSlaveDb() *sqlx.DB {
	return randomBalance()
}

var curIndex int

//roundRobinBalance 轮询获取从数据库
func roundRobinBalance() *sqlx.DB {
	lens := len(slaveDb)
	if curIndex >= lens {
		curIndex = 0
	}
	curIndex = (curIndex + 1) % lens
	return slaveDb[curIndex]
}

//randomBalance 随机获取从数据库
func randomBalance() *sqlx.DB {
	curIndex := rand.Intn(len(slaveDb))
	return slaveDb[curIndex]
}

// Close 关闭MySQL连接
func Close() {
	masterDb.Close()
	for _, db := range slaveDb {
		db.Close()
	}
}
