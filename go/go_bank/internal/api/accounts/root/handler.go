package root

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"

	"github.com/wley3337/learning/tree/main/go/go_bank/internal/handler"
	t "github.com/wley3337/learning/tree/main/go/go_bank/types"
)

type Handler struct {
	DB *sql.DB
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	st := defaultStore{h.DB}

	switch method := r.Method; method {
	case "GET":
		err = handleGetAccounts(w, r, st)
	case "POST":
		err = handleCreateAccount(w, r, st)
	default:
		err = fmt.Errorf("Method not allowed %s", r.Method)
	}

	if err != nil {
		handler.WriteJSONError(w, http.StatusInternalServerError, err)
	}
}

func handleGetAccounts(w http.ResponseWriter, r *http.Request, st defaultStore) error {
	accounts, err := st.GetAccounts()
	if err != nil {
		log.Println("Error getting accounts")
		return err
	}
	return handler.WriteJSON(w, http.StatusOK, accounts)
}

type CreateAccountRequestBody struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber uuid.UUID `json:"accountNumber"`
	Balance       int64     `json:"balance"`
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request, st defaultStore) error {
	req := new(CreateAccountRequestBody)

	// this needs to parse the JSON from the body
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	lastFour := req.AccountNumber.String()[len(req.AccountNumber.String())-4:]

	log.Println("Creating account from:",
		req.FirstName,
		req.LastName,
		lastFour)

	// created a struct that the create account accepts
	account := new(t.CreateAccountAccount)
	account.AccountNumber = req.AccountNumber
	account.FirstName = req.FirstName
	account.LastName = req.LastName

	// call the create account with that struct
	acc, err := st.CreateAccount(account)

	if err != nil {
		return err
	}

	return handler.WriteJSON(w, http.StatusCreated, acc)
}

func handleTransfer(w http.ResponseWriter, r *http.Request, st defaultStore) error {
	return nil
}
