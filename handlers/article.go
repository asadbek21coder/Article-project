package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/project6/models"
	"github.com/saidamir98/project6/storage"
)

func GetArticleList(c *gin.Context) {
	// TODO - read 'search' query param
	search := ""
	resp := storage.Store.GetArticleList(search)

	c.JSON(http.StatusOK, gin.H{
		"message": "GetArticleList",
		"data":    resp,
	})
}

func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	// TODO - get an article by ID
	c.JSON(http.StatusOK, gin.H{
		"message": "GetArticleByID",
		"id":      id,
	})
}

func CreateArticle(c *gin.Context) {
	var data models.CreateArticleModel
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := storage.Store.CreateArticle(models.Article{
		ID:      "123",
		Content: data.Content,
		Author:  data.Author,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO - create an article

	c.JSON(http.StatusCreated, gin.H{
		"message": "Article has been created",
	})
}

func UpdateArticle(c *gin.Context) {
	var data models.Article
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := storage.Store.UpdateArticle(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Article has been updated",
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	// TODO - delete an article by ID

	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteArticle",
		"id":      id,
	})
}
