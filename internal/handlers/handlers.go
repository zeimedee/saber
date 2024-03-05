package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocket(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		time.Sleep(time.Second)
	}
}
