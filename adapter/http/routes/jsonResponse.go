package routes

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	*gin.Context
}

func NewJsonResponse(
	context *gin.Context,
) *JsonResponse {
	return &JsonResponse{
		Context: context,
	}
}

func (jsonResponse *JsonResponse) ThrowError(
	key string,
	err error,
	statusCode int,
) {
	jsonResponse.Writer.Header().Set("Content-Type", "application/json")
	jsonResponse.Writer.WriteHeader(statusCode)

	jsonResponse.JSON(statusCode, map[string]string{key: err.Error()})
}

func (jsonResponse *JsonResponse) SendJson(
	key string,
	data interface{},
	statusCode int,
) {
	jsonResponse.Writer.Header().Set("Content-Type", "application/json")
	jsonResponse.Writer.WriteHeader(statusCode)

	jsonResponse.JSON(statusCode, map[string]interface{}{key: data})
}
