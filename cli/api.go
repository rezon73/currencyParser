package main

import (
	"currencyParser/controller"
	"currencyParser/service/mainDatabase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var routines *sync.WaitGroup
var server   *http.Server

func main() {
	defer mainDatabase.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	routines = &sync.WaitGroup{}
	routines.Add(1)

	handleApiShutDown(sigChan)

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "30000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/get_symbols", controller.IndexController{}.GetSymbolsHandler)
	mux.HandleFunc("/get_quote", controller.IndexController{}.GetQuoteHandler)

	server = &http.Server{
		Addr:    host + ":" + port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	routines.Wait()
}

func handleApiShutDown(signalChan chan os.Signal) {
	go func() {
		defer routines.Done()

		select {
			case sig := <-signalChan:
				log.Print("Caught signal", map[string]interface{}{
					"signal":    sig,
					"operation": "terminating service",
				})

				mainDatabase.Close()
				err := server.Close()
				if err != nil {
					log.Fatal(err)
				}
		}
	}()
}
