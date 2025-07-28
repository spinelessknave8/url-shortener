package main

import (
	"fmt"
	"net/http"
)

func main() {
	http_router := http.NewServeMux()
	http_router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http_server := http.Server{
		Addr:    ":8080",
		Handler: http_router,
	}
	fmt.Println("Listening on port: 8080")
	http_server.ListenAndServe()

}
