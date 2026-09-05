package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	notes  = map[int]Note{}
	nextId = 1
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/notes", notesHandler)
	mux.HandleFunc("/notes/", getNote)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to notes api!\n")
}

func listNotes(w http.ResponseWriter, r *http.Request) {
	allnotes := make([]Note, 0, len(notes))
	for _, note := range notes {
		allnotes = append(allnotes, note)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(allnotes)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	note.Id = nextId
	nextId++

	notes[note.Id] = note

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listNotes(w, r)
	case http.MethodPost:
		createNote(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getNote(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/notes/")
	if path == "" || path == r.URL.Path {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	note, exist := notes[id]
	if !exist {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	if _, exists := notes[id]; !exists {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	var note Note
	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	note.Id = id
	notes[id] = note
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func patchNotes(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.Atoi(path)
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
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
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
	path := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.Atoi(path)
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
