package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed sp24-cpsc-20000-002.json
var syllabusFS embed.FS

func syllabus(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data, err := fs.ReadFile(syllabusFS, "sp24-cpsc-20000-002.json")
		if err != nil {
			http.Error(w, "Failed to read syllabus file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
