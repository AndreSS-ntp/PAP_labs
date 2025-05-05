package main

import (
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/internal/app"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/rpc", app.RPCHandler)
	log.Println("Server started on port 1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
