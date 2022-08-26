package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mimomomimo/web-service-gin/app/model"
	"github.com/stretchr/testify/assert"
)

func TestController_GetBooks(t *testing.T) {
	tests := []struct {
		name          string
		mockRepositry func(url.Values) ([]model.Book, error)
		wantCode      int
	}{
		{
			name: "return 200",
			mockRepositry: func(_ url.Values) ([]model.Book, error) {
				return []model.Book{}, nil
			},
			wantCode: 200,
		},
		{
			name: "error",
			mockRepositry: func(_ url.Values) ([]model.Book, error) {
				return []model.Book{}, errors.New("forced error to marshal")
			},
			wantCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			ginContext, _ := gin.CreateTestContext(response)
			req, _ := http.NewRequest("GET", "/", nil)
			ginContext.Request = req
			repository := &RepositoryInterfaceMock{
				FetchVolumesFunc: tt.mockRepositry,
			}
			controller := &Controller{
				Repositry: repository,
			}
			controller.GetBooks(ginContext)
			var responseBody map[string]interface{}
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
			assert.Equal(t, tt.wantCode, response.Code)
		})
	}
}
