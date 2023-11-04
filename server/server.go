package server

import (
	"log"
	"net/http"
)

func printWelcome(addr string) {
	log.Printf("Loading http server at %s", addr)
	log.Printf("Loading websocket server at %s/ws/", addr)
}
func ListenAndServe(addr string, handler http.Handler) error {
	printWelcome(addr)
	return http.ListenAndServe(addr, handler)
}
func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
	printWelcome(addr)
	return http.ListenAndServeTLS(addr, certFile, keyFile, handler)
}
