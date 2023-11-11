package websocket

import (
	"encoding/json"
	"log"
)

// As in http HandleFunc we pass a type HandlerFunc func(conn *Conn, messageType int,
//
//	arguments ...string), request,,, her we are passing the connection and message type
//	with arguments
func init() { //func(conn *Conn, messageType int, arguments ...string)
	HandleFunc("echo", func(conn *Conn, messageType int, arguments HandlerArguments) {
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
	HandleFunc("commands", func(conn *Conn, messageType int, arguments HandlerArguments) {
		data, err := json.Marshal(Commands())
		if err != nil {
			log.Printf("main: commands: json marshal: error: %s", err)
			return
		}
		if err := conn.WriteMessage(messageType, data); err != nil {
			log.Printf("main: commands: error: %s", err)
		}
	})

	HandleFunc("register", func(conn *Conn, messageType int, arguments HandlerArguments) {

	})
}

func Commands() []string {
	cmds := []string{}
	for k, _ := range router.handlers {
		cmds = append(cmds, k) // Just returns Echo and commands for routes.go
	}
	return cmds
}
