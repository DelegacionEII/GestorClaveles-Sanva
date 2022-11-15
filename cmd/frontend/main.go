package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

const port string = ":3000"

func main() {
	log.Printf("Frontend service started in port: %s", port)

	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until we receive our signal.
	<-ch

	// Revise the shutdown functionality
	// Maybe use gorilla/mux locally here
	/*
		// Create a deadline to wait for
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline
		srv.Shutdown(ctx)
	*/

	log.Println("Frontend service ended")
}
