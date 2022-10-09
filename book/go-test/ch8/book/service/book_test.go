package service

import (
	"context"
	"self/golang-study/book/go-test/ch8/book/model"
	"self/golang-study/book/go-test/ch8/book/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_bookService_CreateBook(t *testing.T) {
	db := util.CreateTestDB(t)
	defer db.Close()

	bookService := NewBookService(db)

	book := &model.Book{
		Title:    "The Go Programming Language",
		AuthorID: 1,
		ISBN:     "978-01111111",
		Subject:  "computers",
	}
	err := bookService.CreateBook(context.TODO(), book)

	require.NoError(t, err)
	assert.NotZero(t, book, book.ID)
}
