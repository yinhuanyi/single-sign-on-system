/**
 * @Author：Robby
 * @Date：2022/1/29 11:52
 * @Function：
 **/

package main

import (
	"gopkg.in/boj/redistore.v1"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store *redistore.RediStore

func Init() (err error) {

	store, err = redistore.NewRediStore(10, "tcp", ":6379", "", []byte("Iay1B9p4Bi28EVsT*qzF"))
	if err != nil {
		log.Printf("session初始化失败：%s\n", err.Error())
		return
	}
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 20,
		HttpOnly: true,
		Secure:   false,
	}

	return
}

func Get(r *http.Request, name string) (sessionValue string, err error) {

	session, err := store.Get(r, "session_id")
	if err != nil {
		return
	}

	if session.Values[name] != nil {
		sessionValue = session.Values[name].(string)
	}

	return
}

func Set(w http.ResponseWriter, r *http.Request, name string, sessionValue string) (err error) {

	session, err := store.Get(r, "session_id")
	if err != nil {
		return
	}

	session.Values[name] = sessionValue

	err = session.Save(r, w)

	return

}

func Delete(w http.ResponseWriter, r *http.Request, name string) (err error) {

	session, err := store.Get(r, "session_id")
	if err != nil {
		return
	}
	delete(session.Values, name)
	err = session.Save(r, w)
	return
}

func main() {
	Init()
	//store.Get("")
}