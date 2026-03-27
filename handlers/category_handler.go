package handlers

import (
    "net/http"

    "bookstore/models"
    "github.com/gin-gonic/gin"
)

var Categories = []models.Category{}
var CategoryID = 1

func GetCategories(c *gin.Context) {
    c.JSON(http.StatusOK, Categories)
}

func CreateCategory(c *gin.Context) {
    var category models.Category

    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if category.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Name required"})
        return
    }

    category.ID = CategoryID
    CategoryID++
    Categories = append(Categories, category)

    c.JSON(http.StatusCreated, category)
}

///go get -u github.com/gin-gonic/gin