package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Ace struct {
	config *conf
	server *http.Server
}

func (a *Ace) Init() {
	a.server = &http.Server{Addr: *addr}
	a.config = &config
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("cloud watch exporter, visit /metric"))
	})
	http.HandleFunc("/metric", metricsHandler)
}

func (a *Ace) Start() {
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	log.Printf("Aws CloudWatch Exporter Version %s", version)
	log.Println("Quit with CONTROL-C.")
	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}

func (a *Ace) Stop() {
	log.Println("Quit The Aws CloudWatch Exporter")
	if err := a.server.Shutdown(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

func RunForever() {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	ace := &Ace{}
	ace.Init()
	ace.Start()
	<-gracefulStop
	ace.Stop()
}
