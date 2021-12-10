package main

import (
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	//url1 := "http://172.22.41.32:8080/test_service_name/v1/do2"
	//url2 := "http://127.0.0.1:8080/test_service_name/v1/do2"
	//url3 := "http://localhost:8080/test_service_name/v1/do2"
	//go Get(url1,"172.22.41.32")

	WebSocketClient()
	//go Get(url2,"127.0.0.1")
	//go Get(url3,"localhost")
	select {

	}

}
func Get(url string, host string){
	method := "GET"
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzYwMzI4MjIsImlzcyI6ImFwcF9pZF9hIn0.Qbt48atk5sCF1CggvMKB-5H2Hm3AxRd80f5EvVDPGxA")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 200;i++ {
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(host," ",time.Now().Minute(), " ", i, " ", string(body))
		time.Sleep(time.Second)
	}
}

func WebSocketClient() {
	fmt.Println("Client started")
	for {
		conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:8888/v1/websocket")
		if err != nil {
			fmt.Println("Cannot connect: " + err.Error())
			time.Sleep(time.Duration(5) * time.Second)
			continue
		}
		fmt.Println("Connected to server")
		for i := 0; i < 10; i++ {
			randomNumber := strconv.Itoa(rand.Intn(100))
			msg := []byte(randomNumber)
			err = wsutil.WriteClientMessage(conn, ws.OpText, msg)
			if err != nil {
				fmt.Println("Cannot send: " + err.Error())
				continue
			}
			fmt.Println("Client message send with random number " + randomNumber)
			msg, _, err := wsutil.ReadServerData(conn)
			if err != nil {
				fmt.Println("Cannot receive data: " + err.Error())
				continue
			}
			fmt.Println("Server message received with random number: " + string(msg))
			time.Sleep(time.Duration(5) * time.Second)
		}
		err = conn.Close()
		if err != nil {
			fmt.Println("Cannot close the connection: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Disconnected from server")
	}
}