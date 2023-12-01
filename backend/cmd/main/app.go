package main

import (
	"backend/internal/config"
	"backend/internal/user"
	"backend/pkg/database"
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net"
	"net/http"
	"time"
)

func main() {
	cfg := config.GetConfig()
	db := database.Init(cfg)

	router := httprouter.New()

	handler := user.NewHandler(db, cfg)
	handler.Register(router)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			_ = fmt.Errorf("can not close database")
		}
	}(db)

	corsHandler := cors.Default().Handler(router)
	start(corsHandler, cfg)
}

func start(handler http.Handler, cfg *config.Config) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
	if err != nil {
		_ = fmt.Errorf("listener was not created")
		panic(err)
	}

	server := &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
