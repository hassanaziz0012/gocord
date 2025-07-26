package web

import (
	"fmt"
	"net/http"
)

type IPAddr struct {
	Host string
	Port uint16
}

var server IPAddr = IPAddr{
	Host: "0.0.0.0",
	Port: 8000,
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%s:%d", ip.Host, ip.Port)
}

func StartHttp() {
	http.ListenAndServe(server.String(), nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
