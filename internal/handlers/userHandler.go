package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/pandukamuditha/simple-blog/internal/common"
	"github.com/pandukamuditha/simple-blog/internal/repositories"
)

func RegisterUserHandlers(r *mux.Router, l *common.Logger, db *pgx.Conn) {
	userHandler := UserHandler{
		logger:  l,
		queries: repositories.New(db),
	}

	r.HandleFunc("/{userId}", userHandler.getUser)
	r.HandleFunc("", userHandler.createUser).Methods(http.MethodPost)
}

type UserHandler struct {
	logger  *common.Logger
	queries *repositories.Queries
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

	val, err := common.StringToInt64(queryParams["userId"])

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := u.queries.GetUser(r.Context(), val)

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

	user, err := u.queries.CreateUser(r.Context(), data)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	common.WriteJsonResponse(rw, http.StatusOK, user)
}
