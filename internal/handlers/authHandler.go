package handlers

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/pandukamuditha/simple-blog/internal/common"
	"github.com/pandukamuditha/simple-blog/internal/repositories"
)

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

	sha_512 := sha512.New()
	sha_512.Write([]byte(loginCredentials.Password))
	inputPasswordHash := hex.EncodeToString(sha_512.Sum(nil))

	if savedPasswordHash != inputPasswordHash {
		common.WriteJsonResponse(
			rw,
			http.StatusBadRequest,
			common.ErrorResponse{Err: "invalid password"},
		)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + (15 * 60000),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "simple-wallet-server",
		Subject:   loginCredentials.Username,
	})

	tokenString, _ := token.SignedString(common.JwtSigningSecret)

	type TokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}

	tokenResponse := TokenResponse{AccessToken: tokenString, TokenType: "Bearer"}
	SendJSONResponse(rw, http.StatusOK, tokenResponse)
}
