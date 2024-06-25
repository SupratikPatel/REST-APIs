

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type BlogEntry struct {
	ID           int    `json:"id"`
	CreationDate string `json:"creationDate"`
	Headline     string `json:"headline"`
	Text         string `json:"text"`
}

var (
	blogEntries []BlogEntry
	mutex       sync.Mutex
	nextID      = 1
)

func getBlogEntries(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	if len(blogEntries) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err := json.NewEncoder(w).Encode(blogEntries)
	if err != nil {
		return
	}
}

func createBlogEntry(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var entry BlogEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	entry.ID = nextID
	nextID++
	blogEntries = append(blogEntries, entry)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}

func getBlogEntryByID(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/blog-entries/"))
	if err != nil {
		http.Error(w, "Invalid blog entry ID", http.StatusBadRequest)
		return
	}
	for _, entry := range blogEntries {
		if entry.ID == id {
			json.NewEncoder(w).Encode(entry)
			return
		}
	}
	http.NotFound(w, r)
}

func updateBlogEntry(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/blog-entries/"))
	if err != nil {
		http.Error(w, "Invalid blog entry ID", http.StatusBadRequest)
		return
	}
	var updatedEntry BlogEntry
	if err := json.NewDecoder(r.Body).Decode(&updatedEntry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, entry := range blogEntries {
		if entry.ID == id {
			updatedEntry.ID = id // Preserve the ID of the blog entry
			blogEntries[i] = updatedEntry
			json.NewEncoder(w).Encode(updatedEntry)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteBlogEntry(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/blog-entries/"))
	if err != nil {
		http.Error(w, "Invalid blog entry ID", http.StatusBadRequest)
		return
	}
	for i, entry := range blogEntries {
		if entry.ID == id {
			blogEntries = append(blogEntries[:i], blogEntries[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/blog-entries", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/blog-entries" {
			http.NotFound(w, r)
			return
		}
		switch r.Method {
		case http.MethodGet:
			getBlogEntries(w, r)
		case http.MethodPost:
			createBlogEntry(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/blog-entries/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/blog-entries/" {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			getBlogEntryByID(w, r)
		case http.MethodPut:
			updateBlogEntry(w, r)
		case http.MethodDelete:
			deleteBlogEntry(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
