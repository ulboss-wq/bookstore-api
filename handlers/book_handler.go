package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "bookstore/models"
)

var Books = []models.Book{}
var BookID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
    category := r.URL.Query().Get("category")
    pageStr := r.URL.Query().Get("page")

    filtered := Books

    if category != "" {
        var catID int
        for _, c := range Categories {
            if strings.EqualFold(c.Name, category) {
                catID = c.ID
            }
        }

        temp := []models.Book{}
        for _, b := range Books {
            if b.CategoryID == catID {
                temp = append(temp, b)
            }
        }
        filtered = temp
    }

    page := 1
    limit := 5

    if pageStr != "" {
        p, _ := strconv.Atoi(pageStr)
        if p > 0 {
            page = p
        }
    }

    start := (page - 1) * limit
    end := start + limit

    if start > len(filtered) {
        start = len(filtered)
    }
    if end > len(filtered) {
        end = len(filtered)
    }

    json.NewEncoder(w).Encode(filtered[start:end])
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    json.NewDecoder(r.Body).Decode(&book)

    if book.Title == "" || book.Price <= 0 {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    book.ID = BookID
    BookID++
    Books = append(Books, book)

    json.NewEncoder(w).Encode(book)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/books/")
    id, _ := strconv.Atoi(idStr)

    for _, b := range Books {
        if b.ID == id {
            json.NewEncoder(w).Encode(b)
            return
        }
    }

    http.NotFound(w, r)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/books/")
    id, _ := strconv.Atoi(idStr)

    var updated models.Book
    json.NewDecoder(r.Body).Decode(&updated)

    for i, b := range Books {
        if b.ID == id {
            updated.ID = id
            Books[i] = updated
            json.NewEncoder(w).Encode(updated)
            return
        }
    }

    http.NotFound(w, r)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/books/")
    id, _ := strconv.Atoi(idStr)

    for i, b := range Books {
        if b.ID == id {
            Books = append(Books[:i], Books[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.NotFound(w, r)
}