package routes

import (
	"github.com/anaclaraddias/brick/core/port"
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	*gin.Context
	connection port.DatabaseConnectionInterface
}

func NewJsonResponse(
	context *gin.Context,
	connection port.DatabaseConnectionInterface,
) *JsonResponse {
	return &JsonResponse{
		Context:    context,
		connection: connection,
	}
}

func (jsonResponse *JsonResponse) ThrowError(
	key string,
	err error,
	statusCode int,
) {
	jsonResponse.Writer.Header().Set("Content-Type", "application/json")
	jsonResponse.Writer.WriteHeader(statusCode)

	jsonResponse.connection.Close()

	jsonResponse.JSON(statusCode, map[string]string{key: err.Error()})
}

func (jsonResponse *JsonResponse) SendJson(
	key string,
	data interface{},
	statusCode int,
) {
	jsonResponse.Writer.Header().Set("Content-Type", "application/json")
	jsonResponse.Writer.WriteHeader(statusCode)

	jsonResponse.connection.Close()

	jsonResponse.JSON(statusCode, map[string]interface{}{key: data})
}
