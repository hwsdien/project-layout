package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"project-layout/internal/config"
	"sync"
	"time"
)

type Mysql struct {
	DB *sqlx.DB
}

var once sync.Once
var db *Mysql

func GetDB() (*Mysql, error) {
	var err error
	once.Do(func() {
		conf := config.GetConfig()
		db = &Mysql{}
		db.DB, err = sqlx.Connect("mysql", conf.Viper.GetString("mysql.dsn"))
		if err != nil {
			fmt.Printf("failed to connect mysql, err:%v\n", err)
			return
		}
		// 设置最大打开的连接数，默认值为0表示不限制。控制应用于数据库建立连接的数量，避免过多连接压垮数据库。
		db.DB.SetMaxOpenConns(conf.Viper.GetInt("mysql.max_open_conns"))
		// 连接池里面允许Idel的最大连接数, 这些Idel的连接 就是并发时可以同时获取的连接,也是用完后放回池里面的互用的连接, 从而提升性能。
		db.DB.SetMaxIdleConns(conf.Viper.GetInt("mysql.max_idle_conns"))
		// 设置一个连接的最长生命周期，因为数据库本身对连接有一个超时时间的设置，如果超时时间到了数据库会单方面断掉连接，
		// 此时再用连接池内的连接进行访问就会出错, 因此这个值往往要小于数据库本身的连接超时时间
		db.DB.SetConnMaxLifetime(time.Duration(conf.Viper.GetInt("mysql.max_life_time")) * time.Second)

		fmt.Printf("ping")
		_ = db.DB.Ping()
	})
	return db, err
}
