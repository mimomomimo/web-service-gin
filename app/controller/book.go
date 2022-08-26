package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mimomomimo/web-service-gin/app/repository"
)

type Controller struct {
	Repositry repository.GoogleBooksRepositry
}

func NewController(r repository.GoogleBooksRepositry) Controller {
	return Controller{
		Repositry: r,
	}
}

func (controller *Controller) GetBooks(c *gin.Context) {
	// リクエスト
	params := c.Request.URL.Query()
	books, err := controller.Repositry.FetchVolumes(params)
	if err != nil {
		// handle error
	}
	// レスポンス
	c.IndentedJSON(http.StatusOK, books)
}
