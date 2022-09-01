package repository

import (
	"net/url"

	"github.com/mimomomimo/web-service-gin/app/model"
)

type RecommendBooksRepositry struct{}

func NewRecommendBooksRepositry() RecommendBooksRepositry {
	return RecommendBooksRepositry{}
}

func (repositry *RecommendBooksRepositry) FetchVolumes(v url.Values) ([]model.Book, error) {
	books := []model.Book{
		{ID: "aaaaa", Title: "オススメの本1", Authors: []string{"著者1"}},
		{ID: "bbbbb", Title: "オススメの本2", Authors: []string{"著者2", "著者3"}},
	}
	return books, nil
}
