package main

import (
	"fmt"
	"geektime/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"os"
)

func newApp(gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID("1"),
		kratos.Name("mydemo"),
		kratos.Version("1.0"),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			gs,
		),
	)
}

func main() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	c := config.New(
		config.WithSource(
			file.NewSource("F:/github/geektime/week3/mydemo/configs/config.json"),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.MydemoConfig
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	//err:=initAPP(bc.Server,bc.Database)
	//if err != nil {
	//	panic(err)
	//}

}
