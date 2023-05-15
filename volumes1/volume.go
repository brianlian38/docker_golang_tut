package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/\n")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
