package main

import (
	"fmt"
	"log"
	"net/http"

	appdb "github.com/thechauhanabhay/blabberit/internal/db"
	"github.com/thechauhanabhay/blabberit/internal/message"
	"github.com/thechauhanabhay/blabberit/internal/user"
)

func main() {
	appdb.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "BlabberIt server is running 🚀")
	})

	http.HandleFunc("/register", user.RegisterHandler)
	http.HandleFunc("/login", user.LoginHandler)

	http.HandleFunc("/send", message.SendMessageHandler)
	http.HandleFunc("/inbox", message.FetchMessagesHandler)

	fmt.Println("Starting BlabberIt on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal((err))
	}
}
