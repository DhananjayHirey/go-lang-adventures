package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Note struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Text      string        `bson:"text" json:"text"`
	Embedding []float32     `bson:"embedding,omitempty" json:"-"`
}

type CreateNoteRequest struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type SearchRequest struct {
	Query string `json:"query"`
}

type SearchResult struct {
	Note  `bson:",inline"`
	Score float64 `bson:"score" json:"score"`
}
