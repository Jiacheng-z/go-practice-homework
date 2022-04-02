package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var serverShutdown = make(chan struct{})

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	server()
	ser := http.Server{Addr: ":8080"}

	g.Go(func() error {
		log.Println("Server Start :8080")
		return ser.ListenAndServe()
	})

	g.Go(func() error {
		<-ctx.Done()
		log.Println("Server Shutdown")

		timeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		return ser.Shutdown(timeout)
	})

	// 退出
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-serverShutdown:
			return errors.New("server request close")
		case <-quit:
			return errors.New("os signal exit")
		}
	})

	if err := g.Wait(); err != nil {
		log.Println("main: error exit ", err.Error())
	}
}

func server() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello!"))
	})
	http.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Shutdown"))
		serverShutdown <- struct{}{}
	})
}
