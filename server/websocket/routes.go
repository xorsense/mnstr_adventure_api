package websocket

import (
	"encoding/json"
	"log"
)

// As in http HandleFunc we pass a type HandlerFunc func(conn *Conn, messageType int,
//  arguments ...string), request,,, her we are passing the connection and message type
//  with arguments
func init() { //func(conn *Conn, messageType int, arguments ...string)
	HandleFunc("echo", func(conn *Conn, messageType int, arguments ...string) {
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

	// Calls a second time
	HandleFunc("commands", func(conn *Conn, messageType int, arguments ...string) {
		data, err := json.Marshal(Commands())
		if err != nil {
			log.Printf("main: commands: json marshal: error: %s", err)
			return
		}
		if err := conn.WriteMessage(messageType, data); err != nil {
			log.Printf("main: commands: error: %s", err)
		}
	})
}
