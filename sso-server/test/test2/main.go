/**
 * @Author：Robby
 * @Date：2022/1/29 21:15
 * @Function：
 **/

package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

// 创建一个基于本地文件系统存储session的store对象，./是session存储路径
var (
	store = sessions.NewFilesystemStore("/Users/Robby/Yinhuanyi_Project/Golang/sso-server/test/test2/session", securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))
)


func set(w http.ResponseWriter, r *http.Request) {

	// 获取名称为user的session对象，如果不存在，会创建一个新的session对象
	session, _ := store.Get(r, "user")

	// 给session设置数据
	session.Values["name"] = "Robby"
	session.Values["age"] = 29

	// 将session对象的用户数据保存到对应的存储中，同时在协议头设置set-cookie，将{user: sessionId}写入浏览器的application中
	err := sessions.Save(r, w)


	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 协议请求
	fmt.Fprintln(w, "Hello World")
}


func get(w http.ResponseWriter, r *http.Request) {

	// 从请求头的cookie字段获取sessionId值，并且从文件系统中加载sessionId对应的session对象
	session, _ := store.Get(r, "user")

	fmt.Fprintf(w, "name:%s age:%d\n", session.Values["name"], session.Values["age"])

}


func main() {

	http.HandleFunc("/setsession", set)
	http.HandleFunc("/getsession", get)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
