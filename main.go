package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pandukamuditha/learn-golang/handlers"
)

func main() {
	logger := log.New(os.Stdout, "posts-api", log.LstdFlags)

	postsHandler := handlers.NewPostsHandler(logger)
	commentsHandler := handlers.NewCommentsHandler(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/posts", postsHandler)
	serveMux.Handle("/comments", commentsHandler)

	httpServer := http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Println("Starting server on 8080")

		err := httpServer.ListenAndServe()

		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	httpServer.Shutdown(ctx)
}
