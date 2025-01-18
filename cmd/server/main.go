package main

import (
	"fmt"
	"net/http"

	"github.com/infysumanta/go-hello-world-server/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.HelloWorldHandler)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}