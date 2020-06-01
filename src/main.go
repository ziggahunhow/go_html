package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("server running on port 5050")
	http.ListenAndServe(":5050", nil)
}
