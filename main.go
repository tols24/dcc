package main

import (
	"Dockerproj/configs"
	"fmt"
	"log"
	"net/http"
)

// alt + t

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Login Page!")
}

func main() { // localhost:8080/testing
	configs.ConnectToDB()
	db := configs.GetDB()
	db.AutoMigrate(Customers{})
	log.Println("Server started on: http://localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Information Page!")
		fmt.Fprintln(w, "End Point 1 -> localhost:8080/login!")
		fmt.Fprintln(w, "End Point 2 -> localhost:8080/logout!")
	})
	http.HandleFunc("/login", Login)

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Logout Page!")
	})

	http.ListenAndServe(":8080", nil)
}
