package handlers

import (
	"net/http"
	"html/template"
	"path/filepath"
	"go-cat-facts/service"
	"go/build"
)

var (
	srcRoot = "/src"
	packageDir = "/go-cat-facts"
	templatePath = build.Default.GOPATH + srcRoot + packageDir + "/templates/"
	indexFile = filepath.Join(templatePath, "index.html")
)

// serve main page
func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	fact, err := service.GetFact()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles(indexFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	context := struct {
		Fact string
	}{
		fact,
	}

	if err := t.Execute(w, context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}