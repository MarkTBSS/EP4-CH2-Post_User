package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "AnuchitO", Age: 18},
}

func usersHandler(write http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		log.Println("GET")
		byte_array, err := json.Marshal(users)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(write, "error: %v", err)
			return
		}
		write.Header().Set("Content-Type", "application/json")
		write.Write(byte_array)
	}
	if request.Method == "POST" {
		log.Println("POST")
		body_byte_array, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Fprintf(write, "error : %v", err)
			return
		}
		new_user := User{}
		err = json.Unmarshal(body_byte_array, &new_user)
		if err != nil {
			fmt.Fprintf(write, "error: %v", err)
			return
		}
		users = append(users, new_user)
		fmt.Printf("% #v\n", users)
		fmt.Fprintf(write, "hello %s created users", "POST")
		return
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)
	log.Println("Server started at :2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
	log.Println("bye bye!")
}
