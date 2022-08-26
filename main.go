package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mimomomimo/web-service-gin/app/controller"
	"github.com/mimomomimo/web-service-gin/app/repository"
)

func main() {
	router := gin.Default()
	//repository := repository.NewGoogleBooksRepositry() // google books apiから書籍情報取得
	repository := repository.NewRecommendBooksRepositry() //固定の書籍情報を取得

	controller := controller.NewController(&repository)
	router.GET("/books", controller.GetBooks)
	router.Run("localhost:8000")
}
