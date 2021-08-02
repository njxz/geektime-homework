package main

import (
	"context"
	"fmt"
	"geektime/week2"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	g.Go(func()error{
		select {
		case sig := <-s:
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("start shutdown server ", sig)
				cancel()
				return errors.New("signal Closed Server")
			}
		}
	})
	g.Go(func() error {
		return week2.Job1(ctx)
	})
	g.Go(func() error {
		return week2.Job2(ctx)
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("quit success")
	}
}
