package main

import (
	"context"
	"log"

	"Semantic-Notes-App/internal/api"
	"Semantic-Notes-App/internal/config"
	"Semantic-Notes-App/internal/database"
	"Semantic-Notes-App/internal/embedding"
	"Semantic-Notes-App/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	ctx := context.Background()

	cfg := config.Load()

	client, err := database.Connect(
		cfg.MongoURI,
	)

	if err != nil {

		log.Fatal(err)

	}

	defer func() {

		err := client.Disconnect(
			context.Background(),
		)

		if err != nil {

			log.Println(err)

		}

	}()

	embedder, err := embedding.NewClient(
		ctx,
		cfg.EmbeddingModel,
		cfg.EmbeddingDimensions,
	)

	if err != nil {

		log.Fatal(err)

	}

	noteRepository := repository.NewNoteRepository(
		client,
		cfg.DatabaseName,
	)

	dims := cfg.EmbeddingDimensions
	if dims == 0 {
		dims = 3072 // gemini-embedding-2 native output size
	}

	// Best-effort: create the Atlas Vector Search index on startup.
	// No-op if it already exists; requires MongoDB Atlas.
	if err := noteRepository.EnsureVectorIndex(ctx, cfg.VectorIndexName, dims); err != nil {
		log.Printf("warning: failed to ensure vector index: %v", err)
	}

	handler := api.NewHandler(
		noteRepository,
		embedder,
		cfg,
	)

	router := gin.Default()

	api.RegisterRoutes(
		router,
		handler,
	)

	err = router.Run(
		":" + cfg.Port,
	)

	if err != nil {

		log.Fatal(err)

	}

}
