package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/HomeProjectSandbox/basicSetup/basicwebserver/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const keyServerAddr = "serverAddr"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.GetRoot)
	mux.HandleFunc("/hello", handler.GetHello)

	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/test", handler.GetTest)

	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()
}
