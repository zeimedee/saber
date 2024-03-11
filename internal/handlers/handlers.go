package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/IBM/sarama"
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

// func (v *valueHandler) AddValue(ctx *gin.Context) {
// 	var body requestBody
// 	if err := ctx.ShouldBindJSON(&body); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	data := requestBody{
// 		Value: body.Value,
// 	}

// 	v.valueService.AddTotal(data.Value)

// 	ctx.JSON(200, gin.H{"message": "value add is : " + strconv.Itoa(v.valueService.Value)})
// }

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (v *valueHandler) WebSocket(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "error creating consumer"})
		return
	}
	defer consumer.Close()
	// partition, err := consumer.Partitions("count")
	// if err != nil {
	// 	ctx.JSON(400, gin.H{"message": "error reading partitions"})
	// 	return
	// }
	pc, err := consumer.ConsumePartition("count", 0, sarama.OffsetNewest)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "error creating partition consumer"})
		return
	}
	defer pc.Close()

	for {
		select {
		case msg := <-pc.Messages():
			val, _ := strconv.Atoi(string(msg.Value))
			v.valueService.AddTotal(val)
			str := strconv.Itoa(v.valueService.Value)
			conn.WriteMessage(websocket.TextMessage, []byte(str))
			time.Sleep(time.Second)
		case err := <-pc.Errors():
			fmt.Printf("Error consuming partition: %v\n", err)
		}

	}
}

func SendValue(ctx *gin.Context) {
	var body requestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := requestBody{
		Value: body.Value,
	}
	services.Producer(data.Value)

	ctx.JSON(200, gin.H{"message": "value is : " + strconv.Itoa(data.Value)})
}
