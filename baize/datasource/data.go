package datasource

import (
	"fmt"
	"github.com/bzdanny/BaiZe/app/setting"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"

	"math/rand"
)

// ProviderSet is datasource providers.
var ProviderSet = wire.NewSet(NewData)

var redisDb *redis.Client

func GetRedisClient() *redis.Client {
	return redisDb

}

// Data .
type Data struct {
	masterDb *sqlx.DB
	slaveDb  []*sqlx.DB
}

// NewData .
func NewData(data *setting.Datasource) (*Data, func(), error) {
	masterDb := newMasterDB(data.Master)
	slaveDb := newSlaveDB(data.Slave)
	client := newRedis(data.Redis)
	cleanup := func() {
		masterDb.Close()
		for _, db := range slaveDb {
			db.Close()
		}
		client.Close()
	}
	redisDb = client
	return &Data{masterDb: masterDb, slaveDb: slaveDb}, cleanup, nil
}

func newMasterDB(master *setting.Master) *sqlx.DB {
	var err error
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", master.User, master.Password, master.Host, master.Port, master.DB)
	masterDb, err := sqlx.Connect(master.DriverName, dsn)
	if err != nil {
		panic(err)
	}
	masterDb.SetMaxOpenConns(master.MaxOpenConns)
	masterDb.SetMaxIdleConns(master.MaxIdleConns)
	return masterDb
}
func newSlaveDB(slave *setting.Slave) []*sqlx.DB {
	count := slave.Count
	var slaveDb []*sqlx.DB
	if count > 0 {
		slaveDb = make([]*sqlx.DB, count)
		var err error
		for i := 0; i < count; i++ {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", slave.Users[i], slave.Passwords[i], slave.Hosts[i], slave.Ports[i], slave.DBs[i])
			slaveDb[i], err = sqlx.Connect(slave.DriverName, dsn)
			if err != nil {
				panic(err)
			}
			slaveDb[i].SetMaxOpenConns(slave.MaxOpenConns)
			slaveDb[i].SetMaxIdleConns(slave.MaxIdleConns)
		}
	}
	return slaveDb
}

func newRedis(r *setting.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
	})
	return rdb
}

// GetMasterDb 获取主数据源
func (d *Data) GetMasterDb() *sqlx.DB {
	return d.masterDb
}

// GetSlaveDb 获取从数据源
func (d *Data) GetSlaveDb() *sqlx.DB {
	return d.roundRobinBalance()
}

var curIndex int

//roundRobinBalance 轮询获取从数据库
func (d *Data) roundRobinBalance() *sqlx.DB {
	lens := len(d.slaveDb)
	if curIndex >= lens {
		curIndex = 0
	}
	curIndex = (curIndex + 1) % lens
	return d.slaveDb[curIndex]
}

//randomBalance 随机获取从数据库
func (d *Data) randomBalance() *sqlx.DB {
	curIndex := rand.Intn(len(d.slaveDb))
	return d.slaveDb[curIndex]
}
