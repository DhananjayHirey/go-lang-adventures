package api

import (
	"net/http"

	"Semantic-Notes-App/internal/config"
	"Semantic-Notes-App/internal/embedding"
	"Semantic-Notes-App/internal/models"
	"Semantic-Notes-App/internal/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository *repository.NoteRepository
	embedder   *embedding.Client
	cfg        *config.Config
}

func NewHandler(
	repository *repository.NoteRepository,
	embedder *embedding.Client,
	cfg *config.Config,
) *Handler {

	return &Handler{
		repository: repository,
		embedder:   embedder,
		cfg:        cfg,
	}

}

func (h *Handler) CreateNote(
	c *gin.Context,
) {

	var request models.CreateNoteRequest

	if err := c.ShouldBindJSON(
		&request,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return

	}

	vector, err := h.embedder.Embed(
		c.Request.Context(),
		request.Title+"\n"+request.Text,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "failed to generate embedding: " + err.Error(),
			},
		)

		return

	}

	note := models.Note{

		Title: request.Title,

		Text: request.Text,

		Embedding: vector,
	}

	err = h.repository.CreateNote(note)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return

	}

	c.JSON(
		http.StatusCreated,
		note,
	)

}

func (h *Handler) GetNotes(
	c *gin.Context,
) {

	notes, err := h.repository.GetNotes()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return

	}

	c.JSON(
		http.StatusOK,
		notes,
	)

}

func (h *Handler) SearchNotes(
	c *gin.Context,
) {

	var request models.SearchRequest

	if err := c.ShouldBindJSON(
		&request,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return

	}

	ctx := c.Request.Context()

	vector, err := h.embedder.Embed(ctx, request.Query)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "failed to generate embedding: " + err.Error(),
			},
		)

		return

	}

	results, err := h.repository.SearchNotes(
		ctx,
		h.cfg.VectorIndexName,
		vector,
		h.cfg.SearchResultLimit,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return

	}

	c.JSON(
		http.StatusOK,
		results,
	)

}
