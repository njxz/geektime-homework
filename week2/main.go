package main

import (
	"context"
	"fmt"
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
	go func() {
		select {
		case sig := <-s:
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("start shutdown server ", sig)
				cancel()
			}
		}
	}()

	g.Go(func() error {
		return Job1(ctx)
	})
	job2 := &Service{Ctx: ctx}

	g.Go(func() error {
		return job2.Start()
	})
	g.Go(func() error {
		return job2.Stop()
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("quit success")
	}
}
