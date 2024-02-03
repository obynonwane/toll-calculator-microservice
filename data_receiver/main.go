package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/obynonwane/tolling/types"
)

// this is a struct of DataReceiver
// with a msg field is a channel
// designed to handle data of type OBUData
type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
}

// create a new instance of the DataReceiver struct
func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
	}
}

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handlerWS) //websocket endpoint
	http.ListenAndServe(":3000", nil)      //serve the server
}

// this function handles the websocket connection
func (dr *DataReceiver) handlerWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	dr.conn = conn

	go dr.wsReceiveLoop()
}

// this func loops through the data
func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("NEW OBU connected client connected!")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}

		fmt.Printf("received OBU data from [%d] :: <lat %.2f, long $\n", data.OBUID)
		//pipe the data into the message channel
		dr.msgch <- data
	}
}
