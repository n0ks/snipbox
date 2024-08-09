package main

import (
  "html/template"
  "path/filepath"
	"github.com/n0ks/snipbox/pkg/models"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template{}) {

  cache := map[string]*template.Template{}
  
}
