package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID    string `json:"id"`
	EMAIL string `json:"email"`
	NAME  string `json:"name"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie User
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	users = append(users, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Conten-type", "application/json")
	//params
	params := mux.Vars(r)
	//loopover the users range
	for index, item := range users {
		if item.ID == params["id"] {
			//delete the users with the i.d. that you've sent
			users = append(users[:index], users[index+1:]...)
			//add a new movie: the movie sent through  posytman
			var movie User
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			users = append(users, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)            //type of parse for the received data
	for index, item := range users { //for index and item in range of users (similar to python)
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...) //append n+1 to end -> 1 to n thus skipping the current value
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	fmt.Println("Hello")
	r := mux.NewRouter()

	users = append(users, User{ID: "1", EMAIL: "one@gmail.com", NAME: "heheheh"})
	users = append(users, User{ID: "2", EMAIL: "two@gmail.com", NAME: "datebayo"})

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Printf("The server starting at port 8010\n")
	log.Fatal(http.ListenAndServe(":8010", r))

}
