/**
 * @Author: Robby
 * @File name: mysql.go
 * @Create date: 2021-05-18
 * @Function:
 **/

package mysqlconnect

import (
	"fmt"
	"goods-server/settings"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Init(cfg *settings.MysqlConfig) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)

	defer func() {
		r := recover()
		if r != nil {
			log.Printf("database connect failed: %s\n", r)
			return
		}
	}()

	// 如果连接不上，这里会panic
	Db = sqlx.MustConnect("mysql", dsn)
	Db.SetMaxOpenConns(cfg.MaxOpenConns)
	Db.SetMaxIdleConns(cfg.MaxIdleConns)
	return

}

func Close() {
	_ = Db.Close()
}
