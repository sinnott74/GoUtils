package util

import "database/sql"

// TxFunc is a function which takes a transaction and returns an error
type TxFunc func(*sql.Tx) error

// Transact wraps a function in a transation
func Transact(db *sql.DB, txFunc TxFunc) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	err = txFunc(tx)
	return err
}
