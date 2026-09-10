package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	notes  = map[int]Note{}
	nextID = 1
)

func main() {
	r := chi.NewRouter()

	r.Route("/notes", func(r chi.Router) {
		r.Get("/", listNotes)
		r.Post("/", createNote)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", getNote)
			r.Put("/", updateNote)
			r.Patch("/", patchNote)
			r.Delete("/", deleteNote)
		})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server starting on :8080")
	log.Fatal(server.ListenAndServe())
}

func listNotes(w http.ResponseWriter, r *http.Request) {
	allNotes := make([]Note, 0, len(notes))
	for _, note := range notes {
		allNotes = append(allNotes, note)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allNotes)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	note.ID = nextID
	nextID++
	notes[note.ID] = note

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	note, exists := notes[id]
	if !exists {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	if _, exists := notes[id]; !exists {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	note.ID = id
	notes[id] = note

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func patchNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	existing, exists := notes[id]
	if !exists {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	var updates map[string]any
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if title, ok := updates["title"]; ok {
		existing.Title = title.(string)
	}
	if content, ok := updates["content"]; ok {
		existing.Content = content.(string)
	}

	notes[id] = existing

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existing)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	if _, exists := notes[id]; !exists {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	delete(notes, id)
	w.WriteHeader(http.StatusNoContent)
}
