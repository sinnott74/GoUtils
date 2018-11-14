package util

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestTransaction(t *testing.T) {
	db, mock, _ := sqlmock.New()

	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectCommit()

	Transact(db, func(tx *sql.Tx) error {
		return nil
	})
}

func TestTransactionError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectRollback()

	txErr := errors.New("Transaction error")

	err := Transact(db, func(tx *sql.Tx) error {
		return txErr
	})

	require.Equal(t, txErr, err)
}

func TestTransactionPanic(t *testing.T) {
	db, mock, _ := sqlmock.New()

	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectRollback()

	txErrPanic := errors.New("Transaction panic")

	defer func() {
		p := recover()
		require.Equal(t, p, txErrPanic)
	}()

	Transact(db, func(tx *sql.Tx) error {
		panic(txErrPanic)
	})
}
