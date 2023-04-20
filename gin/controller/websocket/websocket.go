package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"math/rand"
	"strconv"
)

type WebSocket struct {

}

func (w *WebSocket) Get(ctx *gin.Context){
	//协议升级
	conn, _, _, err := ws.UpgradeHTTP(ctx.Request,ctx.Writer)
	if err != nil {
		log.Default().Println("WebSocket UpgradeHTTP err ", err)
		return
	}
	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			fmt.Println("Error receiving data: " + err.Error())
			fmt.Println("Client disconnected")
			return
		}
		fmt.Println("Client message received with random number: " + string(msg))
		randomNumber := strconv.Itoa(rand.Intn(100))
		err = wsutil.WriteServerMessage(conn, op, []byte(randomNumber))
		if err != nil {
			fmt.Println("Error sending data: " + err.Error())
			fmt.Println("Client disconnected")
			return
		}
		fmt.Println("Server message send with random number " + randomNumber)
	}
}
func Register(eng *gin.RouterGroup) {
	jc := &WebSocket{}
	eng.GET("/v1/websocket", jc.Get)
}