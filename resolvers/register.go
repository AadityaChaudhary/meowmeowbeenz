package resolvers

import (
	"encoding/json"
	"log"
	"net/http"
	"socialcredit/db"
)

type registerParams struct {
	Name    string
	Picture string
}

type registerReturn struct {
	Score float32
	ID    int
}

func register(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		log.Println("not post")
		return
	}
	var p registerParams
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("register", "picture string", p.Picture, "name", p.Name)

	user := db.NewUser(p.Picture, p.Name)
	if user == nil {
		http.Error(w, "dont use the system name", http.StatusBadRequest)
		return
	}

	// create user
	db.DB.Create(user)
	// give user starting points
	db.DB.Create(&db.Rating{Subject: db.SystemID, Object: user.ID, Score: 3})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
