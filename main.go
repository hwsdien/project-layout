package main

import (
	"fmt"
	"project-layout/cmd"
	"project-layout/internal/config"
	"project-layout/internal/db/mysql"
	"project-layout/internal/log"
)

type User struct {
	Id       int64  `db:"u_id"`
	Username string `db:"u_username"`
	Address  string `db:"u_address"`
}

func main() {
	cmd.Execute()

	conf := config.GetConfig()

	fmt.Println(conf.Viper.Get("app_name"))
	fmt.Println(conf.Viper.Get("version"))

	fmt.Println()
	fmt.Println("write log")
	logger := log.GetLogger()
	logger.Log.Info("saldfaadfljljljaljdfjsdaf")
	logger.Log.Error("dasldjasljdflasjdfa")

	db, err := mysql.GetDB()
	if err != nil {
		panic(err)
	}

	users := []User{}
	sql := "select u_id, u_username, u_address from users limit 10"
	if err = db.DB.Select(&users, sql); err == nil {
		for _, user := range users {
			fmt.Println(user.Id, user.Username, user.Address)
		}
	} else {
		panic(err)
	}

}
