package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/pandukamuditha/learn-golang/handlers"
)

func main() {
	logger := log.New(os.Stdout, "blog-api ", log.LstdFlags)

	postsHandler := handlers.NewPostsHandler(logger)
	commentsHandler := handlers.NewCommentsHandler(logger)

	router := mux.NewRouter()
	router.Handle("/posts", postsHandler)
	router.Handle("/comments", commentsHandler)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\": \"ok\"}"))
	})

	httpServer := http.Server{
		Addr:         ":8080",
		Handler:      router,
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

	// Trap interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// Wait 30 seconds and shutdown hhtp server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)

	logger.Print("Shutting down server")
	os.Exit(0)
}
