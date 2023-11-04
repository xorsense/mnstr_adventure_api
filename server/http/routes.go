package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xorsense/mnstr_adventure_api/server/websocket"
)

func init() {
	http.HandleFunc("/ws/", websocket.HandleConnection)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if wrote, err := fmt.Fprintln(w, "Hello, fellow mnstr!"); err != nil {
			log.Printf("main: /: response write: error: %s: wrote: %d", err, wrote)
		}
	})
}
