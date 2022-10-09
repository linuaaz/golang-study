package api

import (
	"context"
	"encoding/json"
	"net/http"
	"self/golang-study/book/go-test/ch8/book/model"
	"self/golang-study/book/go-test/ch8/book/service"
	"strconv"
)

type BookController struct {
	bookService service.BookService
}

func (bc *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	authorID := r.FormValue("authorID")
	isbn := r.FormValue("isbn")
	computers := r.FormValue("subject")

	aid, err := strconv.Atoi(authorID)
	if err != nil {
		http.Error(w, "wrong author id", http.StatusBadRequest)
		return
	}

	book := &model.Book{
		Title:    title,
		AuthorID: aid,
		ISBN:     isbn,
		Subject:  computers,
	}

	err = bc.bookService.CreateBook(context.TODO(), book)
	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(book)

	w.Header().Set("Contetx-Type", "application/json")
	w.Write(data)
}
