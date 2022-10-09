package dao

import (
	"context"
	"self/golang-study/book/go-test/ch8/book/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_bookStore_InsertBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "unexpected err: %s", err)
	defer db.Close()

	store := &bookStore{
		db: db,
	}

	book := &model.Book{
		Title:    "The Go Programming Language",
		AuthorID: 1,
		ISBN:     "978-01111111",
		Subject:  "computers",
	}

	mock.ExpectExec("INSERT INTO books").WillReturnResult(sqlmock.NewResult(1, 1))

	err = store.InsertBook(context.TODO(), book)

	require.NoError(t, err)
	assert.NotZero(t, book.ID)
}
