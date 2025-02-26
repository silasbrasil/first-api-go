package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/silasbrasil/first-api-go/models"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCreateBook_Success(t *testing.T) {
	r := setupRouter()
	createBookRoute(r)

	newBook := models.Book{
		Title:  "Some title",
		Author: "Fulano de Tal",
	}
	bookJson, _ := json.Marshal(newBook)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", strings.NewReader(string(bookJson)))
	r.ServeHTTP(w, req)

	var response models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, newBook.Title, response.Title)
	assert.Equal(t, newBook.Author, response.Author)
}
