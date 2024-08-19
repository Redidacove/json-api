package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var (
	InvalidUser = apiError{Err: "Invalid User", Status: http.StatusForbidden}
	UsersByID = make(map[int]User)
)

type apiError struct{
	Err string
	Status int
}

func (e apiError) Error() string{
	return e.Err
}

type User struct{
	Name  string `json:"Name"`
	ID    int `json:"ID"`
	valid bool `json:"valid"`
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandler(f apiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, apiError{Err: "Internal Server Error", Status: http.StatusInternalServerError})
		}
	}
}
func main(){
	http.HandleFunc("POST /user", makeHTTPHandler(handleCreateUserByID))
	http.HandleFunc("GET /user/{id}", makeHTTPHandler(handleGetUserById))

	http.ListenAndServe(":3000", nil)
}

func handleGetUserById(w http.ResponseWriter, r *http.Request) error{
	if r.Method != http.MethodGet {
		return WriteJSON(w, http.StatusMethodNotAllowed, apiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed})
	}
	id := r.PathValue("id")
	i, err := strconv.Atoi(id)
	if err != nil {
        fmt.Println("Error:", err)
    }

	user := UsersByID[i]
	if !user.valid {
		return WriteJSON(w, http.StatusNotFound, InvalidUser)
	}

	return WriteJSON(w, http.StatusOK, user)
}

func handleCreateUserByID(w http.ResponseWriter, r *http.Request) error{
	if r.Method != http.MethodPost {
		return WriteJSON(w, http.StatusMethodNotAllowed, apiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed})
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	
	if err != nil {
		fmt.Println("Error:", err)
	}

	if user.valid {
			return WriteJSON(w, http.StatusConflict, apiError{Err: "User Already Exists", Status: http.StatusConflict})
	}		
		
	user.valid = true
	UsersByID[user.ID] = user

	return WriteJSON(w, http.StatusCreated, user)
}

func WriteJSON(w http.ResponseWriter, status int , v any) error{
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}