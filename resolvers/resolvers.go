package resolvers

import (
	"fmt"
	"net/http"
)

func Start() {
	fmt.Println("http server starting up")

	http.HandleFunc("/profile", profile)
	http.HandleFunc("/get-scores", getScores)
	http.HandleFunc("/rate", rate)
	http.HandleFunc("/register", register)

	http.ListenAndServe(":8080", nil)
}
