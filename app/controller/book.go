package controller

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/mimomomimo/web-service-gin/app/model"
)

type RepositoryInterface interface {
	FetchVolumes(url.Values) ([]model.Book, error)
}
type Controller struct {
	Repositry RepositoryInterface
}

func NewController(r RepositoryInterface) Controller {
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
