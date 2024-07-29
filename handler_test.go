package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World") // Fprint: Untuk ngeprint ke writer
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/**
 * Handler: Untuk menerima HTTP request yang masuk ke server
 * Handler adalah interface
 * 	Salah satu implementasi HandlerFunc: Membuat function handler HTTP
 */
