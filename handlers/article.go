package handlers

import (
	"net/http"

	"project6/models"

	"github.com/gin-gonic/gin"
)

// ShowAccount GetArticleList
// @ID           get-article-list
// @Summary      Get Article List
// @Description  Get Article List based on query params
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        search  query     string                                             false  "input search text"
// @Param        offset  query     string                                             false  "offset"
// @Param        limit   query     string                                             false  "limit"
// @Success      200     {object}  models.DefaultResponse{data=[]models.ArticleList}  "Success Response"
// @Success      500     {object}  models.DefaultResponse                             "Internal Server Error Response"
// @Router       /articles [GET]
func (h *HandlerImpl) GetArticleList(c *gin.Context) {
	search := c.Query("search")

	offset, err := h.parseOffsetQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limit, err := h.parseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.strg.Article().GetArticleList(models.QueryParams{
		Search: search,
		Offset: offset,
		Limit:  limit,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetArticleList",
		"data":    resp,
	})
}

// ShowAccount GetArticleByID
// @ID           get-article_by_id
// @Summary      Get an article
// @Description  Delete an article based on given id
// @Tags         article
// @Accept       json
// @Produce      json
// @Param id path string true "id"
// @Success      201      {object}  models.DefaultResponse     "Success Response"
// @Success      400      {object}  models.DefaultResponse     "Bad Request Response"
// @Success      500      {object}  models.DefaultResponse     "Internal Server Error Response"
// @Router       /articles/{id} [GET]
func (h *HandlerImpl) GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.strg.Article().GetArticleByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "GetArticleByID",
		"id":      resp,
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
func (h *HandlerImpl) CreateArticle(c *gin.Context) {
	var data models.CreateArticleModel
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.strg.Article().CreateArticle(models.CreateArticleModel{
		Content:  data.Content,
		AuthorID: data.AuthorID,
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

// ShowAccount UpdateArticle
// @ID           update-article
// @Summary      Update an article
// @Description  Ureate an article
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        article  body      models.UpdateArticleModel  true       "article body"
// @Success      201      {object}  models.DefaultResponse     "Success Response"
// @Success      400      {object}  models.DefaultResponse     "Bad Request Response"
// @Success      500      {object}  models.DefaultResponse     "Internal Server Error Response"
// @Router       /articles [PUT]
func (h *HandlerImpl) UpdateArticle(c *gin.Context) {
	var data models.UpdateArticleModel
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.strg.Article().UpdateArticle(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Article has been updated",
	})
}

// ShowAccount DeleteArticle
// @ID           delete-article
// @Summary      Delete an article
// @Description  Delete an article based on given id
// @Tags         article
// @Accept       json
// @Produce      json
// @Param id path string true "id"
// @Success      201      {object}  models.DefaultResponse     "Success Response"
// @Success      400      {object}  models.DefaultResponse     "Bad Request Response"
// @Success      500      {object}  models.DefaultResponse     "Internal Server Error Response"
// @Router       /articles/{id} [DELETE]
func (h *HandlerImpl) DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	err := h.strg.Article().DeleteArticle(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteArticle",
		"id":      id,
	})
}
