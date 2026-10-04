package repository

import (
	"context"
	"strings"

	"Semantic-Notes-App/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type NoteRepository struct {
	collection *mongo.Collection
}

func NewNoteRepository(
	client *mongo.Client,
	databaseName string,
) *NoteRepository {

	collection := client.
		Database(databaseName).
		Collection("notes")

	return &NoteRepository{
		collection: collection,
	}

}

func (r *NoteRepository) CreateNote(
	note models.Note,
) error {

	_, err := r.collection.InsertOne(
		context.Background(),
		note,
	)

	return err

}

func (r *NoteRepository) GetNotes() ([]models.Note, error) {

	cursor, err := r.collection.Find(
		context.Background(),
		struct{}{},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var notes []models.Note

	err = cursor.All(
		context.Background(),
		&notes,
	)

	if err != nil {
		return nil, err
	}

	return notes, nil

}

func (r *NoteRepository) SearchNotes(
	ctx context.Context,
	indexName string,
	queryVector []float32,
	limit int64,
) ([]models.SearchResult, error) {

	numCandidates := limit * 20

	if numCandidates < 100 {
		numCandidates = 100
	}

	pipeline := mongo.Pipeline{

		{{Key: "$vectorSearch", Value: bson.D{
			{Key: "index", Value: indexName},
			{Key: "path", Value: "embedding"},
			{Key: "queryVector", Value: queryVector},
			{Key: "numCandidates", Value: numCandidates},
			{Key: "limit", Value: limit},
		}}},

		{{Key: "$project", Value: bson.D{
			{Key: "embedding", Value: 0},
			{Key: "score", Value: bson.D{
				{Key: "$meta", Value: "vectorSearchScore"},
			}},
		}}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var results []models.SearchResult

	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil

}

type vectorIndexField struct {
	Type          string `bson:"type"`
	Path          string `bson:"path"`
	NumDimensions int32  `bson:"numDimensions"`
	Similarity    string `bson:"similarity"`
}

type vectorIndexDefinition struct {
	Fields []vectorIndexField `bson:"fields"`
}

func (r *NoteRepository) EnsureVectorIndex(
	ctx context.Context,
	indexName string,
	dimensions int32,
) error {

	definition := vectorIndexDefinition{
		Fields: []vectorIndexField{
			{
				Type:          "vector",
				Path:          "embedding",
				NumDimensions: dimensions,
				Similarity:    "cosine",
			},
		},
	}

	model := mongo.SearchIndexModel{
		Definition: definition,
		Options:    options.SearchIndexes().SetName(indexName).SetType("vectorSearch"),
	}

	_, err := r.collection.SearchIndexes().CreateOne(ctx, model)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") || strings.Contains(err.Error(), "already exists") {
			return nil
		}
		return err
	}

	return nil

}
