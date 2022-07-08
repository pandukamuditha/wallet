package handlers

import (
	"log"
	"net/http"
)

type CommentsHandler struct {
	l *log.Logger
}

func NewCommentsHandler(l *log.Logger) *CommentsHandler {
	return &CommentsHandler{l}
}

func (c *CommentsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.getComments(rw, r)
	}
}

func (c *CommentsHandler) getComments(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusAccepted)
}
