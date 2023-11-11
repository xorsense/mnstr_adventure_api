package http

import (
	"fmt"
	"github.com/xorsense/mnstr_adventure_api/server/cors"
	"log"
	"net/http"

	"github.com/xorsense/mnstr_adventure_api/server/websocket"
)

func init() {
	http.HandleFunc("/ws/", cors.CORS(websocket.HandleConnection))
	http.HandleFunc("/", cors.CORS(func(w http.ResponseWriter, r *http.Request) {
		if wrote, err := fmt.Fprintln(w, "Hello, fellow mnstr!"); err != nil {
			log.Printf("main: /: response write: error: %s: wrote: %d", err, wrote)
		}
	}))
}
