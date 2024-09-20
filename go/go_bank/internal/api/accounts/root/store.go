package root

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	t "github.com/wley3337/learning/tree/main/go/go_bank/types"
)

type Store interface {
	GetAccountByID(int) (t.Account, error)
	DeleteAccount(int) error
}

type defaultStore struct {
	db *sql.DB
}

func (st defaultStore) CreateAccount(acc *t.CreateAccountAccount) (t.Account, error) {
	acct := t.Account{}
	query := `
    INSERT INTO account (
        first_name,
        last_name,
        account_number,
        balance
    ) VALUES ($1, $2, $3, $4)
    RETURNING id, first_name, last_name, account_number, balance, created_at, updated_at
    `
	rows, err := st.db.Query(query, acc.FirstName, acc.LastName, acc.AccountNumber, acc.Balance)

	if err != nil {
		return acct, err
	}

	defer rows.Close()
	for rows.Next() {
		err = acct.FromRow(rows)
		return acct, err
	}

	return acct, fmt.Errorf("Failed to create an account")
}

func (st defaultStore) GetAccounts() ([]t.Account, error) {
	query := `
        SELECT
            id,
            first_name,
            last_name,
            account_number,
            balance,
            created_at,
            updated_at
        FROM account
    `
	rows, err := st.db.Query(query)
	if err != nil {
		return nil, err
	}

	accounts := []t.Account{}

	defer rows.Close()
	for rows.Next() {
		acct := t.Account{}
		err := acct.FromRow(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acct)
	}
	return accounts, nil
}
