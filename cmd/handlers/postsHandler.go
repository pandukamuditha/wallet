package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pandukamuditha/simple-blog/cmd/common"
)

func RegisterPostsHandlers(r *mux.Router, l *common.Logger) {
	postHandler := NewPostHandler(l)

	r.HandleFunc("/posts", postHandler.getPostsHandler).Methods("GET")

	postsSubRouter := r.PathPrefix("/posts").Subrouter()
	postsSubRouter.HandleFunc("/{id}", postHandler.getPostHandler).Methods("GET")
}

type PostsHandler struct {
	logger *common.Logger
}

func NewPostHandler(l *common.Logger) *PostsHandler {
	return &PostsHandler{l}
}

func (p *PostsHandler) getPostsHandler(rw http.ResponseWriter, r *http.Request) {
	p.logger.Log("Getting a list of posts")
	rw.WriteHeader(http.StatusOK)
}

func (p *PostsHandler) getPostHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p.logger.Logf("Getting post: %s", vars["id"])
	rw.WriteHeader(http.StatusOK)
}
