package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type volumes struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID         string     `json:"id"`
	VolumeInfo volumeInfo `json:"volumeInfo"`
}

type volumeInfo struct {
	Title string `json:"title"`
}

func main() {
	router := gin.Default()
	repositry := NewGoogleBooksRepositry()
	controller := NewController(repositry)
	router.GET("/books", controller.getBooks)
	router.Run("localhost:8000")
}

type Controller struct {
	Repositry GoogleBooksRepositry
}

func NewController(r GoogleBooksRepositry) Controller {
	return Controller{
		Repositry: r,
	}
}

func (controller *Controller) getBooks(c *gin.Context) {
	params := c.Request.URL.Query()
	books, err := controller.Repositry.FetchVolumes(params)
	if err != nil {
		// handle error
	}
	c.IndentedJSON(http.StatusOK, books)
}

type GoogleBooksRepositry struct {
	URL string
}

func NewGoogleBooksRepositry() GoogleBooksRepositry {
	url := "https://www.googleapis.com/books/v1/volumes"
	return GoogleBooksRepositry{URL: url}
}

func (repositry *GoogleBooksRepositry) FetchVolumes(v url.Values) ([]book, error) {
	resp, err := http.Get(repositry.URL + "?" + v.Encode())
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var volumes volumes
	if err := json.Unmarshal(body, &volumes); err != nil {
		// handle error
	}

	var books []book
	for _, v := range volumes.Items {
		book := book{
			ID:    v.ID,
			Title: v.VolumeInfo.Title,
		}
		books = append(books, book)
	}
	return books, err
}
