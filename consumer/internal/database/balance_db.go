package database

import (
	"database/sql"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (a *BalanceDB) GetBalance(accountId string) (float64, error) {
	stmt, err := a.DB.Prepare("SELECT amount FROM balance WHERE account_id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var balance float64
	row := stmt.QueryRow(accountId)
	err = row.Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (a *BalanceDB) UpdateBalance(accountId string, balance float64) error {
	stmtDelete, err := a.DB.Prepare("DELETE FROM balance WHERE account_id = ?")
	if err != nil {
		return err
	}
	defer stmtDelete.Close()
	_, err = stmtDelete.Exec(accountId)
	if err != nil {
		return err
	}

	stmtInsert, err := a.DB.Prepare("INSERT INTO balance(account_id, amount) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmtInsert.Close()
	_, err = stmtInsert.Exec(accountId, balance)
	if err != nil {
		return err
	}
	return nil
}
