package main

import (
	"github.com/gohouse/gorose"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"./config"
)

func main() {
	DB := gorose.Connect.Open(config.DbConfig, "mysql_dev")
	// close DB
	defer DB.Close()
	// get the db chaining object
	var db gorose.Database

	user := db.Table("users a").
		Join("area b","a.id","=","b.uid").
		Where("id",">",1).Get()

	fmt.Println(db.LastSql())
	fmt.Println(user)

	// return json
	//res2 := user.Limit(2).Get()
	//fmt.Println(db.LastSql())
	//fmt.Println(user)

}

