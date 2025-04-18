package main

import (
	"fmt"
	"net/http"

	"github.com/thechauhanabhay/blabberit/internal/user"
)

func main() {
	user.InitDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "BlabberIt server is running 🚀")
	})

	http.HandleFunc("/register", user.RegisterHandler)

	fmt.Println("Starting BlabberIt on :8080...")
	http.ListenAndServe(":8080", nil)
}
