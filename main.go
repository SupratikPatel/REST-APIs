/*package main

import "fmt"

var age = 22
var name = "SuperUser"
var name2 string

const (
	one = iota
	two
	three
)
const four = 4

var b [3]int = [3]int{1, 2, 3}

//a:= [3]int{1,2,3}
//c:=[7]int{1,2,3,4,5,6}
//s:=c[2:5]
//l:=len(s)
//s:=[]int{1, 2, 3} slices

func main() {
	fmt.Println("Hello StackIT")
	name2 := "xyz"
	c := [7]int{1, 2, 3, 4, 5, 6}
	s := c[2:5]
	l := len(s)
	d := cap(s)
	fmt.Println(l)
	fmt.Println(d)
	fmt.Println("My name is " + name + name2)
}
*/
/*
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

// BlogEntry represents the structure of a blog entry.
type BlogEntry struct {
	ID           int    `json:"id"`
	CreationDate string `json:"creationDate"`
	Headline     string `json:"headline"`
	Text         string `json:"text"`
}

var (
	blogEntries []BlogEntry
	mutex       sync.Mutex
)

// getBlogEntries handles GET requests to fetch all blog entries.
func getBlogEntries(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	if len(blogEntries) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(blogEntries)
}

// createBlogEntry handles POST requests to create a new blog entry.
func createBlogEntry(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var entry BlogEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	blogEntries = append(blogEntries, entry)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}

// getBlogEntryByID handles GET requests to fetch a single blog entry by ID.
func getBlogEntryByID(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	idStr := strings.TrimPrefix(r.URL.Path, "/blog-entries/")
	id, err := strconv.Atoi(idStr)
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

// updateBlogEntry handles PUT requests to update an existing blog entry.
func updateBlogEntry(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	idStr := strings.TrimPrefix(r.URL.Path, "/blog-entries/")
	id, err := strconv.Atoi(idStr)
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
			blogEntries[i] = updatedEntry
			json.NewEncoder(w).Encode(updatedEntry)
			return
		}
	}
	http.NotFound(w, r)
}

// deleteBlogEntry handles DELETE requests to remove a blog entry.
func deleteBlogEntry(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	idStr := strings.TrimPrefix(r.URL.Path, "/blog-entries/")
	id, err := strconv.Atoi(idStr)
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

	// Register handlers for the "/blog-entries" path
	mux.HandleFunc("/blog-entries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBlogEntries(w, r)
		case http.MethodPost:
			createBlogEntry(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Register handlers for the "/blog-entries/{id}" path
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
*/

/*
GET Request (Retrieve All Blog Entries)
http
GET http://localhost:8080/blog-entries
Accept: application/json

POST Request (Create a New Blog Entry)
http
POST http://localhost:8080/blog-entries
Content-Type: application/json

{
  "creationDate": "2024-04-25",
  "headline": "My First Blog Post",
  "text": "This is the content of my first blog post."
}

GET Request (Retrieve a Single Blog Entry by ID)
Replace 1 with the actual ID of the blog entry you want to retrieve.
http
GET http://localhost:8080/blog-entries/1
Accept: application/json

PUT Request (Update an Existing Blog Entry)
Replace 1 with the actual ID of the blog entry you want to update.
http
PUT http://localhost:8080/blog-entries/1
Content-Type: application/json

{
  "id": 1,
  "creationDate": "2024-04-26",
  "headline": "Updated Blog Post",
  "text": "This is the updated content of the blog post."
}

DELETE Request (Delete a Blog Entry)
Replace 1 with the actual ID of the blog entry you want to delete.
http
DELETE http://localhost:8080/blog-entries/1
*/

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
