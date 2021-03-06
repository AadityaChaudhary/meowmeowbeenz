package resolvers

import (
	"encoding/json"
	"log"
	"net/http"
	"socialcredit/db"
)

type rateParams struct {
	SubjectID uint
	ObjectID  uint
	Score     int
}

func rate(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		log.Println("not post")
		return
	}
	var p rateParams
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("subjectID", p.SubjectID, "objectID", p.ObjectID, "score", p.Score)

	if p.ObjectID == 0 || p.SubjectID == 0 {
		http.Error(w, "missing ID", http.StatusBadRequest)
		return
	}
	if p.Score < 1 || p.Score > 5 {
		http.Error(w, "score should be between 1 and 5", http.StatusBadRequest)
		return
	}

	// check if the rating exists
	rating := db.Rating{Subject: p.SubjectID, Object: p.ObjectID, Score: p.Score}

	if db.DB.Model(&rating).Where("subject = ?", p.SubjectID).Where("object = ?", p.ObjectID).Update("score", p.Score).RowsAffected == 0 {
		log.Println("rating doesnt exist")
		// create user
		db.DB.Create(&rating)
	}
	log.Println("after executed update")

	// trigger recount
	var object db.User
	db.DB.First(&object, p.ObjectID)
	score, err := db.RecountScore(object)
	if err != nil {
		http.Error(w, "missing ID", http.StatusBadRequest)
		return
	}
	object.Score = score

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&object)

}
