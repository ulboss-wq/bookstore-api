package handlers

import (
    "net/http"
    "strings"
    "strconv"

    "bookstore/models"
    "github.com/gin-gonic/gin"
)

var Books = []models.Book{}
var BookID = 1

func GetBooks(c *gin.Context) {
    category := c.Query("category")
    filtered := Books

    if category != "" {
        var catID int
        for _, cat := range Categories {
            if strings.EqualFold(cat.Name, category) {
                catID = cat.ID
            }
        }

        temp := []models.Book{}
        for _, b := range filtered {
            if b.CategoryID == catID {
                temp = append(temp, b)
            }
        }
        filtered = temp
    }

    pageStr := c.DefaultQuery("page", "1") 
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
        return
    }

    limit := 10 
    start := (page - 1) * limit
    end := start + limit
    if start > len(filtered) {
        start = len(filtered)
    }
    if end > len(filtered) {
        end = len(filtered)
    }

    paged := filtered[start:end]

    c.JSON(http.StatusOK, gin.H{
        "page":       page,
        "limit":      limit,
        "total":      len(filtered),
        "totalPages": (len(filtered) + limit - 1) / limit,
        "books":      paged,
    })
}

func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if book.Title == "" || book.Price <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    book.ID = BookID
    BookID++
    Books = append(Books, book)
    c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, b := range Books {
        if b.ID == id {
            c.JSON(http.StatusOK, b)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func UpdateBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var updated models.Book
    if err := c.ShouldBindJSON(&updated); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, b := range Books {
        if b.ID == id {
            updated.ID = id
            Books[i] = updated
            c.JSON(http.StatusOK, updated)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, b := range Books {
        if b.ID == id {
            Books = append(Books[:i], Books[i+1:]...)
            c.Status(http.StatusNoContent)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}