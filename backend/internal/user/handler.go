package user

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	// для админ панели нужен POST, GET, UPDATE, PATCH
	TestUrl = "/test"
	// для абитуры только GET??
)

type handler struct {
	db  *sql.DB
	cfg *config.Config
}

func NewHandler(db *sql.DB, cfg *config.Config) handlers.Handler {
	return &handler{db, cfg}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(TestUrl, h.Test)
}

func (h *handler) Test(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from backend!"))
}
