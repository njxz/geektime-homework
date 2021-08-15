//+build !wireinject

package main

import (
	"fmt"
	"geektime/internal/conf"
	"geektime/internal/data"
)

func initAPP(server *conf.MyDemoServer, database *conf.MydemoDatabase) error {
	db, _, err := data.NewData(database)
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
	return nil
}
