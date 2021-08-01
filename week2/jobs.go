package week2

import (
	"context"
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
func Job2(ctx context.Context) error {
	http.HandleFunc("/", handle)
	srv := http.Server{Addr: ":8081"}
	go func() {
		<-ctx.Done()
		srv.Shutdown(ctx)
	}()
	err := srv.ListenAndServe()
	fmt.Println("server shutdown ", err)
	fmt.Println("job2 stop")
	return err
}
