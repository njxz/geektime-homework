package main

import (
	"encoding/json"
	"fmt"
	"geektime/internal/conf"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"os"
)

func decoder(kc *config.KeyValue, v map[string]interface{}) error {
	err := json.Unmarshal(kc.Value, &v)
	fmt.Println(err)
	return nil
}
func main() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	c := config.New(
		config.WithSource(
			file.NewSource("configs/week3_config.yaml"),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}
	fmt.Println(c.Value("server.grpc"))
	var bc conf.MydemoConfig
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	fmt.Println(bc)

}
