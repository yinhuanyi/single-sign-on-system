package session

import (
	"log"
	"net/http"
	"sso/settings"

	"gopkg.in/boj/redistore.v1"

	"github.com/gorilla/sessions"
)

var store *redistore.RediStore

func Init(cfg *settings.SessionConfig) (err error) {

	store, err = redistore.NewRediStore(10, "tcp", ":6379", "", []byte(cfg.HashKey))
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

	session, err := store.Get(r, settings.Conf.SessionConfig.SessionId)
	if err != nil {
		return
	}

	if session.Values[name] != nil {
		sessionValue = session.Values[name].(string)
	}

	return
}

func Set(w http.ResponseWriter, r *http.Request, name string, sessionValue string) (err error) {

	session, err := store.Get(r, settings.Conf.SessionConfig.SessionId)
	if err != nil {
		return
	}

	session.Values[name] = sessionValue

	err = session.Save(r, w)

	return

}

func Delete(w http.ResponseWriter, r *http.Request, name string) (err error) {

	session, err := store.Get(r, settings.Conf.SessionConfig.SessionId)
	if err != nil {
		return
	}
	delete(session.Values, name)
	err = session.Save(r, w)
	return
}
