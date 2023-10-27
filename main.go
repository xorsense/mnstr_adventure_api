package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/xorsense/mnstr_adventure_api/server/websocket"
)

func main() {
	http.HandleFunc("/ws/", websocket.HandleConnection)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if wrote, err := fmt.Fprintln(w, "Hello, fellow mnstr!"); err != nil {
			log.Printf("main: /: response write: error: %s: wrote: %d", err, wrote)
		}
	})

	websocket.HandleFunc("echo", func(conn *websocket.Conn, messageType int, arguments ...string) {
		log.Printf("echo: message type: %d, arguments: %v", messageType, arguments)
		data, err := json.Marshal(arguments)
		if err != nil {
			log.Printf("main: echo: json marshal: error: %s", err)
			return
		}
		if err := conn.WriteMessage(messageType, data); err != nil {
			log.Printf("main: echo: error: %s", err)
		}
	})

	websocket.HandleFunc("commands", func(conn *websocket.Conn, messageType int, arguments ...string) {
		data, err := json.Marshal(websocket.Commands())
		if err != nil {
			log.Printf("main: commands: json marshal: error: %s", err)
			return
		}
		if err := conn.WriteMessage(messageType, data); err != nil {
			log.Printf("main: commands: error: %s", err)
		}
	})

	log.Printf("loading server at loclahost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
