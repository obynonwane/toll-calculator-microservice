package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/obynonwane/tolling/types"
)

func main() {
	fmt.Println("data receiver working fine")
}

// this is a struct of DataReceiver
// with a msg field is a channel
// designed to handle data of type OBUData
type DataReceiver struct {
	msg  chan types.OBUData
	conn *websocket.Conn
}

func (dr DataReceiver) handlerWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1028, 1028)
	if err != nil {
		log.Fatal(err)
	}
}
