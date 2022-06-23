package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/project6/models"
	"github.com/saidamir98/project6/storage"
)

// ShowAccount GetArticleList
// @ID           get-article-list
// @Summary      Get Article List
// @Description  Get Article List based on query params
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        search  query     string                                         false  "input search text"
// @Success      200     {object}  models.DefaultResponse{data=[]models.Article}  "Success Response"
// @Success      500     {object}  models.DefaultResponse                         "Internal Server Error Response"
// @Router       /articles [GET]
func GetArticleList(c *gin.Context) {
	search := c.Query("search")

	resp, err := storage.Store.GetArticleList(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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

// ShowAccount CreateArticle
// @ID           create-article
// @Summary      Create an article
// @Description  Create an article based on given body
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        article  body      models.CreateArticleModel  true  "article body"
// @Success      201      {object}  models.DefaultResponse     "Success Response"
// @Success      400      {object}  models.DefaultResponse     "Bad Request Response"
// @Success      500      {object}  models.DefaultResponse     "Internal Server Error Response"
// @Router       /articles [POST]
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "Article has been created",
		"data":    nil,
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
