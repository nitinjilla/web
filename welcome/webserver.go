// BASIC WEB SERVICE

// Author: Nitin Jilla

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

var port string = ":8000"
var version string = "v1.1"

func main() {

	r := mux.NewRouter()
	srv := http.Server{
		Addr:    port,
		Handler: r,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	r.HandleFunc("/", homePage)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	<-c
	log.Println("Shutting down...")
	srv.Shutdown(context.Background())
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!\nVersion: %s\n ", version)
}
