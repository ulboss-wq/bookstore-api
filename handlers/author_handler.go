package handlers

import (
    "net/http"

    "bookstore/models"
    "github.com/gin-gonic/gin"
)

var Authors = []models.Author{}
var AuthorID = 1

func GetAuthors(c *gin.Context) {
    c.JSON(http.StatusOK, Authors)
}

func CreateAuthor(c *gin.Context) {
    var author models.Author

    if err := c.ShouldBindJSON(&author); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if author.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Name required"})
        return
    }

    author.ID = AuthorID
    AuthorID++
    Authors = append(Authors, author)

    c.JSON(http.StatusCreated, author)
}