package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"

	"github.com/DatTVu/pratice-golang/react_app_ver_2/backend/pkg/config"
	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

//var ch = make(chan bytes.Buffer)

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func parseArgs(s string) []string {
	if len(s) <= 0 {
		return nil
	}
	zp := regexp.MustCompile("[\\s]+")
	return zp.Split(s, -1)
}

func execCommand() {
	cmd := exec.Command("firefox")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if val, ok := config.Commands[string(p)]; ok {
			test := parseArgs(string(p))
			fmt.Println("[Reader]: Parsing Done")
			val().Run(test)
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			//result := <-ch
			log.Println(err)
			return
		}

	}
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	config.InitCommands()
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
