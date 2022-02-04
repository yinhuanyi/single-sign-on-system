/**
 * @Author: Robby
 * @File name: snowflake.go
 * @Create date: 2021-05-19
 * @Function: 使用雪花算法的模块
 **/

package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

// Init 算法参数初始化
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// GenID 生成ID
func GenID() int64 {
	return node.Generate().Int64()
}
