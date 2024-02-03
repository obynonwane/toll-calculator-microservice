package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

// time  interval of every second
var sendInterval = time.Second

// define a websocket endpoint
const wsEndpoint = "ws://127.0.0.1:30000/ws"

// struct defined for GPS data
type OBUData struct {
	OBUID int     `json:"ubuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `josn:"long"`
}

// function to return generated coordinates
func gentLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

// function to generate coordinates
func genCoord() float64 {
	//generate random number between 0 and 100 but nerver zero by adding 1 to it
	n := float64(rand.Intn(100) + 1)
	//generate random float
	f := rand.Float64()

	return n + f
}
func main() {
	obuIDS := generateOBUIDS(100)

	//initialise websocker using default dialer
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	//define an infinite loop
	for {
		for i := 0; i < len(obuIDS); i++ {
			lat, long := gentLatLong()
			data := OBUData{
				OBUID: obuIDS[i],
				Lat:   lat,
				Long:  long,
			}

			fmt.Printf("%+v\n", data)

			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		//sleep for 60 seconds
		time.Sleep(sendInterval)
	}
}

// generate a slice of integer
func generateOBUIDS(n int) []int {
	//creates a slice of integer and initialises it to length n
	ids := make([]int, n)

	// Generate random integers and populate the slice
	for i := 0; i < n; i++ {
		// Generate a random integer in the range [0, math.MaxInt) ensuring a broad range
		ids[i] = rand.Intn(math.MaxInt)
	}

	return ids
}

// init function will call just before your program starts
func init() {
	//sequence of random numbers generated is different each time the program is run
	rand.Seed(time.Now().UnixNano())
}
