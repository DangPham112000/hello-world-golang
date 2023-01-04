package main

import (
	"fmt"
	"net/http"

	"github.com/DangPham112000/hello-world-golang/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting listening on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
