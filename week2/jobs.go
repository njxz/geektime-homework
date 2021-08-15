package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Job1(ctx context.Context) error {
	in := os.Stdin
	out := os.Stdout
	b := make([]byte, 10)
	//fmt.Println(in.Stat())
	for {
		select {
		case <-ctx.Done():
			fmt.Println("job1 stop")
			return nil
		default:
		}
		n, err := io.ReadFull(in, b[:1])
		if err != nil {
			return err
		}
		if string(b[0]) == "a" {
			return errors.New("Job1 errors ")
		}
		_, err = out.Write(b[:n])
		if err != nil {
			return err
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	time.Sleep(time.Second * 3)
	w.Write([]byte("<h1>welcome to geektime!!</h1>"))

}

type Servic interface {
	Start() error
	Stop() error
}
type Service struct {
	Ctx context.Context
	srv http.Server
}

func (s *Service) Start() error {
	http.HandleFunc("/", handle)
	srv := http.Server{Addr: ":8081"}
	err := srv.ListenAndServe()
	fmt.Println("server shutdown ", err)
	fmt.Println("job2 stop")
	return err
}
func (s *Service) Stop() error {
	<-s.Ctx.Done()
	tctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	err := s.srv.Shutdown(tctx)
	return err
}
