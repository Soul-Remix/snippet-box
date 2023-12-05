package main

import "github.com/Soul-Remix/snippet-box/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
