package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type Conn struct {
	*websocket.Conn
}

type HandlerFunc func(conn *Conn, messageType int, arguments ...string)

type Command struct {
	Namespace string   `json:"namespace"`
	Arguments []string `json:"arguments"`
}

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) HandleFunc(namespace string, handler HandlerFunc) {
	if r.handlers == nil {
		r.handlers = make(map[string]HandlerFunc) // just allocs space
	}
	r.handlers[namespace] = handler // instantiates both
}

func (r *Router) Parse(namespace string) (HandlerFunc, error) {
	handler, ok := r.handlers[namespace]
	if !ok {
		return nil, fmt.Errorf("namespace not found: %s", namespace)
	}
	return handler, nil
}

func (r *Router) ParseJSON(data []byte) (*Command, HandlerFunc, error) {
	var command Command
	if err := json.Unmarshal(data, &command); err != nil {
		return nil, nil, err
	}
	handler, err := r.Parse(command.Namespace)
	if err != nil {
		return nil, nil, err
	}
	return &command, handler, nil
}
