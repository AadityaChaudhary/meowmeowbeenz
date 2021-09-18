package resolvers

import (
	"encoding/json"
	"net/http"
	"socialcredit/db"
	"strings"
)

type Object struct {
	Rating  int
	Score   float32
	Picture string
	Name    string
}

func GetScore(w http.ResponseWriter, req *http.Request) {

	var subjectID string
	var objectsString string

	if subjectID = req.URL.Query().Get("subjectID"); subjectID == "" {
		// error
		http.Error(w, "no subject id", http.StatusBadRequest)
		return
	}

	if objectsString = req.URL.Query().Get("objects"); objectsString == "" {
		// error
		http.Error(w, "no object id", http.StatusBadRequest)
		return
	}
	objectIDs := strings.Split(objectsString, ",")
	var objects []Object

	for _, id := range objectIDs {
		var user db.User
		db.DB.First(&user, id)
		var rating db.Rating
		db.DB.Model(&rating).Where("subject = ?", subjectID).Where("object = ?", id).First(&rating)
		objects = append(objects, Object{Rating: rating.Score, Score: user.Score, Picture: user.Picture, Name: user.Name})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&objects)
}
