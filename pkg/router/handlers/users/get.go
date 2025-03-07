package users_handlers

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET ALL USERS"))
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET USER BY ID USERS"))
}
