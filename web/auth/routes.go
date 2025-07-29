package auth

import (
	"fmt"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sign up to Gocord")
}
