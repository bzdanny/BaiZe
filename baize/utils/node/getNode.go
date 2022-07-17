package node

import (
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/utils/ipUtils"
	"github.com/gogf/gf/v2/util/gconv"
	"math/rand"
	"time"
)

func GetNodeId() int {
	rand.Seed(time.Now().UnixNano())
	var id int
	for true {
		id = rand.Intn(1023)
		val := datasource.GetRedisClient().Get("snowflake:" + gconv.String(id)).Val()
		if val == "" {
			break
		}
	}
	s, err := ipUtils.GetLocalIP()
	if err != nil {
		panic(err)
	}

	go func(id string, data string) {
		for true {
			datasource.GetRedisClient().Set("snowflake:"+id, s, time.Hour+time.Minute).Err()
			t := time.NewTimer(time.Hour)
			<-t.C
		}
	}(gconv.String(id), s)
	return id
}
