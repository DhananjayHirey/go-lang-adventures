package embedding

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

// Client wraps the Google GenAI client for generating text embeddings.
type Client struct {
	genaiClient *genai.Client
	model       string
	dimensions  int32
}

// NewClient creates a new embedding client. It relies on genai.NewClient
// auto-detecting the API key from the GEMINI_API_KEY / GOOGLE_API_KEY
// environment variable (already loaded via godotenv in config.Load()).
// model defaults to "gemini-embedding-2" if empty.
// dimensions is optional; pass 0 to use the model's native output size (3072).
func NewClient(ctx context.Context, model string, dimensions int32) (*Client, error) {

	if model == "" {
		model = "gemini-embedding-2"
	}

	genaiClient, err := genai.NewClient(ctx, nil)

	if err != nil {
		return nil, fmt.Errorf("embedding: failed to create genai client: %w", err)
	}

	return &Client{
		genaiClient: genaiClient,
		model:       model,
		dimensions:  dimensions,
	}, nil

}

// Embed converts a single piece of text into an embedding vector.
func (c *Client) Embed(ctx context.Context, text string) ([]float32, error) {

	contents := []*genai.Content{
		genai.NewContentFromText(text, genai.RoleUser),
	}

	var config *genai.EmbedContentConfig

	if c.dimensions > 0 {
		dims := c.dimensions
		config = &genai.EmbedContentConfig{
			OutputDimensionality: &dims,
		}
	}

	result, err := c.genaiClient.Models.EmbedContent(
		ctx,
		c.model,
		contents,
		config,
	)

	if err != nil {
		return nil, fmt.Errorf("embedding: failed to embed content: %w", err)
	}

	if len(result.Embeddings) == 0 || len(result.Embeddings[0].Values) == 0 {
		return nil, fmt.Errorf("embedding: no embedding values returned")
	}

	return result.Embeddings[0].Values, nil

}
