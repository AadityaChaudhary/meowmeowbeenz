package db

import (
	"errors"
	"log"
)

type Score struct {
	Score  float32
	Rating int
}

func RecountScore(object User) (float32, error) {
	// get all scores that have the given objectID
	var ratings []Rating
	DB.Where("object = ?", object.ID).Find(&ratings)
	if len(ratings) < 1 {
		//ruh roh
		return 0, errors.New("no ratings found")
	}
	var total float32
	var scores []Score
	for _, rating := range ratings {
		var u User
		DB.Select("score").First(&u, rating.Subject)
		scores = append(scores, Score{Score: u.Score, Rating: rating.Score})
		total += u.Score
	}
	log.Println("scores", scores)

	var finalScore float32
	for _, score := range scores {
		finalScore += float32(score.Rating) * (score.Score / total)
	}

	DB.Model(&object).Update("score", finalScore)
	return finalScore, nil
}
