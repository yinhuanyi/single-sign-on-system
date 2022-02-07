/**
 * @Author：Robby
 * @Date：2022/2/6 15:03
 * @Function：
 **/

package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	// 加载模型文件和策略文件，获取enforce对象
	e, err := casbin.NewEnforcer("./test/model.conf", "./test/policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "robby1", "/goods/list", "read")
	check(e, "robby1", "/goods/list", "write")
	check(e, "robby2", "/goods/list", "read")
	check(e, "robby2", "/goods/list", "write")
}