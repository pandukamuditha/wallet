package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pandukamuditha/learn-golang/cmd/handlers"
)

func main() {
	logger := log.New(os.Stdout, "blog-api ", log.LstdFlags)

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	} else {
		logger.Print("Successfully loaded .env file")
	}

	appServerPort := os.Getenv("APP_SERVER_PORT")

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
		Addr:         fmt.Sprintf(":%s", appServerPort),
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Printf("Starting server on %s", appServerPort)

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

	// Wait 30 seconds and shutdown http server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)

	logger.Print("Shutting down server")
	os.Exit(0)
}
