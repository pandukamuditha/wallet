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
	"github.com/pandukamuditha/simple-blog/internal/common"
	"github.com/pandukamuditha/simple-blog/internal/handlers"
	"github.com/pandukamuditha/simple-blog/internal/middleware"
)

// @title           Basic Wallet
// @version         1.0
// @description     This is a basic wallet application
// @host      		localhost:8008

func main() {
	logger := common.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Log("Error loading .env file")
	} else {
		logger.Log("Successfully loaded .env file")
	}

	appServerPort := os.Getenv("APP_SERVER_PORT")

	router := mux.NewRouter()

	db := common.Connect(os.Getenv("DB_URL"), logger)

	// health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\": \"ok\"}"))
	})

	// swagger ui
	fs := http.FileServer(http.Dir("./docs"))
	router.PathPrefix("/docs").Handler(http.StripPrefix("/docs", fs))

	// user related routes
	handlers.RegisterUserHandlers(router.PathPrefix("/user").Subrouter(), logger, db)

	// auth related routes
	handlers.RegisterAuthHandlers(router.PathPrefix("/auth").Subrouter(), logger, db)

	// wallet related routes
	handlers.RegisterWalletHandlers(router.PathPrefix("/wallet").Subrouter(), logger, db)

	// token validation only
	router.Use(middleware.AuthenticationMiddleware)

	readTimeout, err := common.GetEnvInt("READ_TIMEOUT")
	if err != nil {
		readTimeout = 5
	}

	writeTimeout, err := common.GetEnvInt("READ_TIMEOUT")
	if err != nil {
		writeTimeout = 5
	}

	idleTimeout, err := common.GetEnvInt("READ_TIMEOUT")
	if err != nil {
		idleTimeout = 5
	}

	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%s", appServerPort),
		Handler:      router,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
	}

	go func() {
		logger.Logf("Starting server on %s", appServerPort)

		err := httpServer.ListenAndServe()

		if err != nil {
			logger.Logf("Error starting server: %s\n", err)
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

	logger.Log("Shutting down server")
	os.Exit(0)
}
