package handlers

import (
	"crypto"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/pandukamuditha/simple-blog/internal/common"
	"github.com/pandukamuditha/simple-blog/internal/repositories"
)

var hmacSampleSecret []byte

func RegisterAuthHandlers(router *mux.Router, logger *common.Logger, db *pgx.Conn) {
	authHandler := AuthHandler{
		logger:  logger,
		queries: *repositories.New(db),
	}

	router.HandleFunc("/login", authHandler.login)
}

type AuthHandler struct {
	logger  *common.Logger
	queries repositories.Queries
}

func (a *AuthHandler) login(rw http.ResponseWriter, r *http.Request) {
	type LoginCredentials struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var loginCredentials LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&loginCredentials)

	if err != nil || loginCredentials.Username == "" || loginCredentials.Password == "" {
		common.WriteJsonResponse(
			rw,
			http.StatusBadRequest,
			common.ErrorResponse{Err: "invalid username/password"},
		)
		return
	}

	savedPasswordHash, err := a.queries.GetPasswordHash(r.Context(), loginCredentials.Username)

	if err != nil {
		common.WriteJsonResponse(
			rw,
			http.StatusBadRequest,
			common.ErrorResponse{Err: "invalid username"},
		)
		return
	}

	sha512 := crypto.SHA512.New()
	sha512.Write([]byte(loginCredentials.Password))
	inputPasswordHash := string(sha512.Sum(nil))

	if savedPasswordHash != inputPasswordHash {
		common.WriteJsonResponse(
			rw,
			http.StatusBadRequest,
			common.ErrorResponse{Err: "invalid password"},
		)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, _ := token.SignedString(hmacSampleSecret)

	type TokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}

	tokenResponse := TokenResponse{AccessToken: tokenString, TokenType: "Bearer"}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(tokenResponse)
}
