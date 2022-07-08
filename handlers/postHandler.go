package handlers

import (
	"log"
	"net/http"
)

type PostsHandler struct {
	l *log.Logger
}

func NewPostsHandler(l *log.Logger) *PostsHandler {
	return &PostsHandler{l}
}

func (p *PostsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getPosts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *PostsHandler) getPosts(rw http.ResponseWriter, r *http.Request) {
	postsList := "Sample Posts List"

	rw.Write([]byte(postsList))
}
