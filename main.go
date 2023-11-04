package main

import (
	"log"

	"github.com/xorsense/mnstr_adventure_api/server"
)

func main() {
	log.Fatal(server.ListenAndServe(":8080", nil))
}
