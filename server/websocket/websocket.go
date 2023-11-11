package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}
var router = NewRouter()

func UpgradeHTTP(w http.ResponseWriter, r *http.Request, h http.Header) (*websocket.Conn, error) {
	return upgrader.Upgrade(w, r, h)
}

func HandleFunc(namespace string, handler HandlerFunc) {
	router.HandleFunc(namespace, handler)
}

func Parse(namespace string) (HandlerFunc, error) {
	return router.Parse(namespace)
}

func ParseJSON(data []byte) (*Command, HandlerFunc, error) {
	return router.ParseJSON(data)
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	log.Printf("websocket: handle connection: upgrading request")
	c, err := UpgradeHTTP(w, r, nil)
	if err != nil {
		log.Printf("websocket: handle connection: upgrade http: error: %s", err)
	}
	conn := &Conn{c}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("websocket: close connection: error: %s", err)
		}
	}()

	// Loop and wait for messages to handle
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("websocket: handle connection: read message: error: %s", err)
			break
		}
		if mt == websocket.BinaryMessage {
			if err := conn.WriteMessage(mt, []byte("binary messages are not supported at this time")); err != nil {
				log.Printf("websocket: handle connection: write message: error: %s", err)
			}
			break
		}
		cmd, handler, err := ParseJSON(msg)
		if err != nil {
			log.Printf("websocket: handle connection: parse json: error: %s", err)
			break
		}
		handler(conn, mt, cmd.Arguments)
	}
}
