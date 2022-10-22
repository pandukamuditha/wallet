package handlers

import (
	"fmt"
	"net/http"

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
	fmt.Print(r.Context().Value(middleware.TokenClaimsKey{}))
}
