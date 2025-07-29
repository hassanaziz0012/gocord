package web

import (
	"gocord/web/auth"
	"net/http"
)

var server IPAddr = IPAddr{
	Host: "0.0.0.0",
	Port: 8000,
}

func StartHttp() {
	http.Handle("/signup", http.HandlerFunc(auth.SignupHandler))
	http.ListenAndServe(server.String(), nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
