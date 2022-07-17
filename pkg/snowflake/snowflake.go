package snowflake

import (
	"github.com/bzdanny/BaiZe/baize/utils/node"
	"github.com/gogf/gf/v2/util/gconv"

	"time"

	sf "github.com/bwmarrin/snowflake"
)

var sfNode *sf.Node

func Init(startTime string) (err error) {
	id := node.GetNodeId()
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	sfNode, err = sf.NewNode(gconv.Int64(id))

	return
}

func GenID() int64 {
	return sfNode.Generate().Int64()
}
