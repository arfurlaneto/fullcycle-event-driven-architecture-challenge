package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("Create table if not exists clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("Create table if not exists accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("Create table if not exists transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("Create table if not exists balance (account_id varchar(255), amount int)")
	if err != nil {
		panic(err)
	}
}
