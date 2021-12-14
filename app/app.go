package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book struct (Model)
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Init books var as a slice Book struct
var books []Book

type App struct {
	Server *mux.Router
}

func (app *App) Initialize() {

	// Mock data
	books = append(books, Book{ID: "1", Title: "Bumi Manusia", Author: "Pramoedya Ananta Toer"})
	books = append(books, Book{ID: "2", Title: "Negeri 5 Menara", Author: "Ahmad Fuadi"})

	app.Server = mux.NewRouter()
	app.Routes()
}

// Handle baseurl request
func (app *App) Handle(w http.ResponseWriter, _ *http.Request) {
	// w.Header().Set("Content-Type", "Application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

// GetBooks Get all books
func (app *App) GetBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

// GetBook Get single book
func (app *App) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r) // Get params
	fmt.Println(r)
	// Loop through the books and find with ID
	for _, item := range books {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}

	err := json.NewEncoder(w).Encode(&Book{})
	if err != nil {
		return
	}
}

// CreateBooks Create Books
func (app *App) CreateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe.
	books = append(books, book)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}
}

// UpdateBook Update book
func (app *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				return
			}
			return
		}
	}

	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

// DeleteBook Delete book
func (app *App) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func (app *App) Routes() {
	app.Server.HandleFunc("/", app.Handle).Methods("GET")
	app.Server.HandleFunc("/api/books", app.GetBooks).Methods("GET")
	app.Server.HandleFunc("/api/books/{id}", app.GetBook).Methods("GET")
	app.Server.HandleFunc("/api/books", app.CreateBooks).Methods("POST")
	app.Server.HandleFunc("/api/books/{id}", app.UpdateBook).Methods("PUT")
	app.Server.HandleFunc("/api/books/{id}", app.DeleteBook).Methods("DELETE")
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":8000", app.Server))
}
