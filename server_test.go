package go_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// Membuat server
	server := http.Server{ // Server adalah struct
		Addr: "localhost:8080", // host:port (Disarankan 4 digit port)
	}

	err := server.ListenAndServe() // Untuk menjalankan server
	if err != nil {
		panic(err)
	}
}

/**
 * Di dalam golang sudah tersedia web server
 * package net/http: Untuk membuat web
 * Server hanya berfungsi sebagai web server
 */
