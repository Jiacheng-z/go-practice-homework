package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	eng := gin.Default()
	sev := initApp(eng)
	sev.RegisterService()

	g.Go(func() error {
		log.Println("Server Start :8080")
		return eng.Run(":8080")
	})

	// 退出
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-quit:
			return errors.New("os signal exit")
		}
	})

	if err := g.Wait(); err != nil {
		log.Println("main: error exit ", err.Error())
	}
}
