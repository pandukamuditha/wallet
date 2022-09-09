package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/pandukamuditha/simple-blog/internal/common"
	"github.com/pandukamuditha/simple-blog/internal/models"
	"github.com/pandukamuditha/simple-blog/internal/repositories"
)

func RegisterUserHandlers(r *mux.Router, l *common.Logger, db *pgx.Conn) {
	userHandler := UserHandler{
		logger: l,
		dao:    &repositories.UserRepository{DB: db},
	}

	r.HandleFunc("/{userId}", userHandler.getUser)
	r.HandleFunc("", userHandler.createUser).Methods(http.MethodPost)
}

type UserDAO interface {
	GetUserById(ctx context.Context, userId int) (*models.User, error)
	CreateUser(ctx context.Context, args repositories.CreateUserParams) (*models.User, error)
}

type UserHandler struct {
	logger *common.Logger
	dao    UserDAO
}

// ShowAccount godoc
// @Summary      Get a user
// @Description  Get user by user ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userId	path int true "User ID"
// @Success      200  {object}  models.User
// @Failure      500
// @Router       /user/{userId} [get]
func (u *UserHandler) getUser(rw http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)

	val, err := strconv.Atoi(queryParams["userId"])

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := u.dao.GetUserById(r.Context(), val)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	json.NewEncoder(rw).Encode(user)
}

func (u *UserHandler) createUser(rw http.ResponseWriter, r *http.Request) {
	var data repositories.CreateUserParams
	json.NewDecoder(r.Body).Decode(&data)

	user, err := u.dao.CreateUser(r.Context(), data)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	json.NewEncoder(rw).Encode(user)
}
