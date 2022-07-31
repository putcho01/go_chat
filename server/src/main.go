package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/putcho01/go_chat/src/domain"
	"github.com/putcho01/go_chat/src/handlers"
)

func main() {
	hub := domain.NewHub()
	go hub.RunLoop()

	http.HandleFunc("/ws", handlers.NewWebsocketHandler(hub).Handle)

	port := "80"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Server Error:", err)
	}
}