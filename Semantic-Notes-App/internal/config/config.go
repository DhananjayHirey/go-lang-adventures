package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	MongoURI            string
	DatabaseName        string
	NotesCollection     string
	GeminiAPIKey        string
	EmbeddingModel      string
	EmbeddingDimensions int32
	VectorIndexName     string
	SearchResultLimit   int64
}

func Load() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	dims, _ := strconv.Atoi(os.Getenv("EMBEDDING_DIMENSIONS"))

	limit, err := strconv.ParseInt(os.Getenv("SEARCH_RESULT_LIMIT"), 10, 64)
	if err != nil || limit <= 0 {
		limit = 5
	}

	cfg := &Config{
		Port:                os.Getenv("PORT"),
		MongoURI:            os.Getenv("MONGO_URI"),
		DatabaseName:        os.Getenv("DATABASE_NAME"),
		NotesCollection:     os.Getenv("NOTES_COLLECTION"),
		GeminiAPIKey:        os.Getenv("GEMINI_API_KEY"),
		EmbeddingModel:      os.Getenv("EMBEDDING_MODEL"),
		EmbeddingDimensions: int32(dims),
		VectorIndexName:     os.Getenv("VECTOR_INDEX_NAME"),
		SearchResultLimit:   limit,
	}

	if cfg.Port == "" {
		log.Fatal("PORT missing")
	}

	if cfg.MongoURI == "" {
		log.Fatal("MONGO_URI missing")
	}

	if cfg.GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY missing")
	}

	if cfg.EmbeddingModel == "" {
		cfg.EmbeddingModel = "text-embedding-004"
	}

	if cfg.VectorIndexName == "" {
		cfg.VectorIndexName = "vector_index"
	}

	return cfg

}
