package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	acctDetail "github.com/wley3337/learning/tree/main/go/go_bank/internal/api/accounts/detail"
	acctRoot "github.com/wley3337/learning/tree/main/go/go_bank/internal/api/accounts/root"
)

// wrapper to allow for handling of errors acts as a decorator
type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func RunServer(addr string, db *sql.DB) error {
	router := mux.NewRouter()
	accountsSubRouter := router.PathPrefix("/api/accounts").Subrouter()
	accountsSubRouter.Handle("/{id}", acctDetail.Handler{DB: db})
	accountsSubRouter.Handle("/", acctRoot.Handler{DB: db})

	// router.Handle("/api/accounts/", accounts.Routes{DB: db})
	// router.Handle("/accounts/{id}", detail.Handler{})
	// router.Handle("/accounts", root.Handler{DB: db})
	log.Println("JSON API server running on port:", addr)
	return http.ListenAndServe(addr, router)
}
