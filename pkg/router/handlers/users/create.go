package users_handlers

import "net/http"

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CREATE USER"))
}
