package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	Port         int    `mapstructure:"port"`
	Host         string `mapstructure:"host"`
	*TokenConfig `mapstructure:"token"`
	*LogConfig   `mapstructure:"log"`
	*Datasource  `mapstructure:"datasource"`
}

type TokenConfig struct {
	ExpireTime int64  `mapstructure:"expire_time"`
	Secret     string `mapstructure:"secret"`
	Issuer     string `mapstructure:"issuer"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type BerbixConfig struct {
	SecretId    string `mapstructure:"secret_id"`
	TemplateKey string `mapstructure:"template_key"`
}

type PlaidConfig struct {
	PlaidClientId string `mapstructure:"plaid_client_id"`
	PlaidSecret   string `mapstructure:"plaid_secret"`
	PlaidEnv      string `mapstructure:"plaid_env"`
}

type Datasource struct {
	Master   *Master   `mapstructure:"master"`
	Slave    *Slave    `mapstructure:"slave"`
	Redis    *Redis    `mapstructure:"redis"`
	RabbitMQ *RabbitMQ `mapstructure:"rabbit_mq"`
}
type Master struct {
	DriverName   string `mapstructure:"driver_name"`
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type Slave struct {
	Count        int      `mapstructure:"count"`
	DriverName   string   `mapstructure:"driver_name"`
	Hosts        []string `mapstructure:"hosts"`
	Users        []string `mapstructure:"users"`
	Passwords    []string `mapstructure:"passwords"`
	DBs          []string `mapstructure:"dbnames"`
	Ports        []int    `mapstructure:"ports"`
	MaxOpenConns int      `mapstructure:"max_open_conns"`
	MaxIdleConns int      `mapstructure:"max_idle_conns"`
}

type Redis struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type RabbitMQ struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func Init(filePath string) {
	// 方式1：直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	//viper.SetConfigFile("./conf/config.yaml")
	// 绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("/Users/xxx/Desktop/bluebell/conf/config.yaml")

	// 方式2：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	// 配置文件名不需要带后缀
	// 配置文件位置可配置多个
	//viper.SetConfigName("config") // 指定配置文件名（不带后缀）
	//viper.AddConfigPath(".") // 指定查找配置文件的路径（这里使用相对路径）
	//viper.AddConfigPath("./conf")      // 指定查找配置文件的路径（这里使用相对路径）

	// 基本上是配合远程配置中心使用的，告诉viper当前的数据使用什么格式去解析
	//viper.SetConfigType("json")

	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		panic(err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
			panic(err)
		}
	})
	return
}
