package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/library/models"
	"github.com/library/utils"
	"github.com/stretchr/testify/assert"
)

func setupData() {
	db, _, ctx, _, err := models.Connect(utils.DbURL)
	if err != nil {
		log.Fatal("Could not setup db")
	}

	var books []interface{}
	for i := 0; i < 5; i++ {
		books = append(books, models.Book{ID: 1, Name: fmt.Sprintf("Test%d", i)})
	}
	db.Collection(utils.Books).InsertMany(ctx, books)
}
func TestGetAllBook(t *testing.T) {
	t.Parallel()

	setupData()
	r, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()

	GetBooks(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetSingleBook(t *testing.T) {
	t.Parallel()

	setupData()
	r, _ := http.NewRequest("GET", "/books/1", nil)
	w := httptest.NewRecorder()

	GetBooks(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestPostSingleBook(t *testing.T) {
	t.Parallel()

	setupData()
	book, _ := json.Marshal(models.Book{ID: 10, Name: "newBook"})
	r, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(book))
	w := httptest.NewRecorder()

	CreateBook(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	book, _ = json.Marshal(models.Book{ID: 10, Name: "newBook"})
	r, _ = http.NewRequest("POST", "/books", bytes.NewBuffer(book))
	w = httptest.NewRecorder()

	CreateBook(w, r)
	assert.Equal(t, http.StatusConflict, w.Code)

}
