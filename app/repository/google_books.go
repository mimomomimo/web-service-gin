package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/mimomomimo/web-service-gin/app/model"
)

// GoogleBooks APIの構造
type Volumes struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID         string     `json:"id"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type VolumeInfo struct {
	Title string `json:"title"`
}

type GoogleBooksRepositry struct {
	URL string
}

// APIのスキーマ：https://developers.google.com/books/docs/v1/reference/volumes/list
func NewGoogleBooksRepositry() GoogleBooksRepositry {
	url := "https://www.googleapis.com/books/v1/volumes"
	return GoogleBooksRepositry{URL: url}
}

func (repositry *GoogleBooksRepositry) FetchVolumes(v url.Values) ([]model.Book, error) {
	//リクエスト処理の参考 https://pkg.go.dev/net/http
	resp, err := http.Get(repositry.URL + "?" + v.Encode())
	if err != nil {
		//handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//handle error
	}

	var volumes Volumes
	// jsonのマッピング
	if err := json.Unmarshal(body, &volumes); err != nil {
		// handle error
	}
	books := volumes.toBooks()
	return books, err
}

func (volumes *Volumes) toBooks() []model.Book {
	var books []model.Book
	for _, v := range volumes.Items {
		book := model.Book{
			ID:    v.ID,
			Title: v.VolumeInfo.Title,
		}
		books = append(books, book)
	}
	return books
}
