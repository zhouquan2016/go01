package main

import (
	"go01/handler"
	"log"
	"net/http"
)

func main() {
	handler.RegisterHandlers()
	log.Fatalln(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil))
}
