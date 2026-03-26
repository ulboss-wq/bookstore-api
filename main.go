package main

import (
    "log"
    "net/http"

    "bookstore/handlers"
)

func main() {

    http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            handlers.GetBooks(w, r)
        } else if r.Method == http.MethodPost {
            handlers.CreateBook(w, r)
        }
    })

    http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            handlers.GetBookByID(w, r)
        case http.MethodPut:
            handlers.UpdateBook(w, r)
        case http.MethodDelete:
            handlers.DeleteBook(w, r)
        }
    })

    http.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            handlers.GetAuthors(w, r)
        } else if r.Method == http.MethodPost {
            handlers.CreateAuthor(w, r)
        }
    })

    http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            handlers.GetCategories(w, r)
        } else if r.Method == http.MethodPost {
            handlers.CreateCategory(w, r)
        }
    })

    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}