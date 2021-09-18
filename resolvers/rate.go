package resolvers

import (
	"encoding/json"
	"net/http"
	"socialcredit/db"

	"gorm.io/gorm"
)

type rateParams struct {
	SubjectID uint
	ObjectID  uint
	Score     int
}

func rate(w http.ResponseWriter, req *http.Request) {
	var p rateParams
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	if err = db.DB.Model(&rating).Where("subject = ?", p.SubjectID).Where("object = ?", p.ObjectID).Update("score", p.Score).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// create user
			db.DB.Create(&rating)
		}
	}
}
