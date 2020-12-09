package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	//服务器开关
	stopC := make(chan struct{})
	//创建带有cancel的ctx
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//创建errgroup
	g, _ := errgroup.WithContext(ctx)

	// 启动server
	server1 := newServer("server1", ":19990")
	server2 := newServer("server2", ":19991")

	g.Go(func() error {
		if err := server1.ListenAndServe(); err != nil {
			cancel()
			return err
		}
		return nil
	})

	g.Go(func() error {
		if err := server2.ListenAndServe(); err != nil {
			cancel()
			return err
		}
		return nil
	})

	//mock error
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		err :=  errors.New("mock error")
		cancel()
		return err
	})

	//Catch signal
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for sig := range sigc {
			needStop := false
			switch sig {
			case syscall.SIGTERM:
				//do something like write log or data backup
				log.Println("SIGTERM catched")
				needStop = true
			case syscall.SIGHUP:
				//do something like write log or data backup
				log.Println("SIGHUP catched")
				needStop = true
			case syscall.SIGINT:
				//do something like write log or data backup
				log.Println("SIGINT catched")
				needStop = true
			default:

			}
			if needStop == true {
				//do something
				cancel()
			}
		}
	}()

	go func() {
		<- ctx.Done()
		go func() {
			log.Println(ctx.Err())
			if err1 := server1.Shutdown(ctx); err1 != nil {
				log.Println("server1 shutdown err: [%v]\n", err1)
			}
			if err2 := server2.Shutdown(ctx); err2 != nil {
				log.Println("server2 shutdown err: [%v]\n", err2)
			}
			log.Println("all servers shutdown")
			close(stopC)
			return
		}()
		// 超时保护
		<-time.After(time.Minute * 5)
		log.Println("shutdown over-time, force to exit")
		close(stopC)
		return
	}()
	<- stopC
}
