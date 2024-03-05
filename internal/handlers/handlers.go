package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zeimedee/saber/internal/services"
)

type requestBody struct {
	Value int `json:"value"`
}
type valueHandler struct {
	valueService *services.Value
}

func NewValueHandler(value *services.Value) *valueHandler {
	return &valueHandler{
		valueService: value,
	}
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "server is happy 😀"})
}

func (v *valueHandler) AddValue(ctx *gin.Context) {
	var body requestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := requestBody{
		Value: body.Value,
	}

	v.valueService.AddTotal(data.Value)

	ctx.JSON(200, gin.H{"message": "value add is : " + strconv.Itoa(v.valueService.Value)})
}

func (v *valueHandler) GetValue(ctx *gin.Context) {
}
