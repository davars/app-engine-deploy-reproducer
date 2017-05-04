package guestbook

import (
	"net/http"

	"mypkgs/handlers"
)

func init() {
	http.HandleFunc("/", handlers.Root)
	http.HandleFunc("/sign", handlers.Sign)
}
