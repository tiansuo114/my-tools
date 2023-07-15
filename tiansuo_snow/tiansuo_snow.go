package tiansuo_snow

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func init() {
	var st time.Time
	err := error(nil)
	startTime := time.Now().Format("2006-01-02")
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	// 设置时间

	machineID := getMachineID()

	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

// 返回int64位的 id值
func Generate_SnowflakeID_i64() int64 {
	return node.Generate().Int64()
}

func getMachineID() int64 {
	return 1000
}

func Generate_SnowflakeID_str() string {
	return node.Generate().String()
}
