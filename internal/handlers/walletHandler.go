package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/pandukamuditha/simple-blog/internal/common"
	"github.com/pandukamuditha/simple-blog/internal/middleware"
	"github.com/pandukamuditha/simple-blog/internal/repositories"
)

func RegisterWalletHandlers(r *mux.Router, l *common.Logger, db *pgx.Conn) {
	walletHandler := WalletHandler{
		logger:  l,
		queries: repositories.New(db),
	}

	r.HandleFunc("/{walletId}/balance", walletHandler.getWalletBalance)
}

type WalletHandler struct {
	logger  *common.Logger
	queries *repositories.Queries
}

func (w *WalletHandler) getWalletBalance(rw http.ResponseWriter, r *http.Request) {
	type WalletBalanceResponse struct {
		WalletId int64 `json:"walletId"`
		Balance  int32 `json:"balance"`
	}

	params := mux.Vars(r)

	walletId, err := common.StringToInt64(params["walletId"])

	if err != nil {
		common.WriteJsonResponse(
			rw,
			http.StatusBadRequest,
			common.ErrorResponse{Err: "invalid wallet id"},
		)
		return
	}

	wallet, err := w.queries.GetWalletById(r.Context(), walletId)

	if err != nil {
		common.WriteJsonResponse(
			rw,
			http.StatusBadRequest,
			common.ErrorResponse{Err: "wallet not found for wallet id"},
		)
		return
	}

	tokenClaims, ok := r.Context().Value(middleware.TokenClaimsKey{}).(jwt.StandardClaims)

	if !ok {
		common.WriteJsonResponse(
			rw,
			http.StatusBadRequest,
			common.ErrorResponse{Err: "error processing auth token"},
		)
		return
	}

	// get userId from the username
	authUserId, err := w.queries.GetUserIdByUsername(r.Context(), tokenClaims.Subject)

	if err != nil {
		common.WriteJsonResponse(
			rw,
			http.StatusInternalServerError,
			common.ErrorResponse{Err: "authenticated user not found"},
		)
		return
	}

	// validate permission to access wallet
	if wallet.UserID != authUserId {
		common.WriteJsonResponse(
			rw,
			http.StatusForbidden,
			common.ErrorResponse{Err: "insuficient permissions to access wallet"},
		)
		return
	}

	walletBalanceResponse := WalletBalanceResponse{
		WalletId: wallet.ID,
		Balance:  wallet.Balance,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(walletBalanceResponse)
}

func (w *WalletHandler) transferToWallet(rw http.ResponseWriter, r *http.Request) {
	type TransferToWalletRequest struct {
		WalletIdFrom int64 `json:"walletIdFrom"`
		WalletIdTo   int64 `json:"walletIdTo"`
		Amount       int32 `json:"amount"`
	}

	var transferToWalletRequest TransferToWalletRequest
	json.NewDecoder(r.Body).Decode(&transferToWalletRequest)

	if (transferToWalletRequest.Amount <= 0) {
		return
	}

	
}
