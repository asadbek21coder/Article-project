package handlers

import (
	"net/http"

	"project6/models"

	"github.com/gin-gonic/gin"
)

// ShowAccount GetAuthorList
// @ID           get-author-list
// @Summary      Get Author List
// @Description  Get Author List based on query params
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        search  query     string                                            false  "input search text"
// @Param        offset  query     string                                            false  "offset"
// @Param        limit   query     string                                            false  "limit"
// @Success      200     {object}  models.DefaultResponse{data=[]models.AuthorList}  "Success Response"
// @Success      500     {object}  models.DefaultResponse                            "Internal Server Error Response"
// @Router       /authors [GET]
func (h *HandlerImpl) GetAuthorList(c *gin.Context) {
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

	resp, err := h.strg.Author().GetAuthorList(models.QueryParams{
		Search: search,
		Offset: offset,
		Limit:  limit,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetAuthorList",
		"data":    resp,
	})
}

func (h *HandlerImpl) GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	// TODO - get an author by ID
	c.JSON(http.StatusOK, gin.H{
		"message": "GetAuthorByID",
		"id":      id,
	})
}

// ShowAccount CreateAuthor
// @ID           create-author
// @Summary      Create an author
// @Description  Create an author based on given body
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        author  body      models.CreateAuthorModel  true  "author body"
// @Success      201     {object}  models.DefaultResponse    "Success Response"
// @Success      400     {object}  models.DefaultResponse    "Bad Request Response"
// @Success      500     {object}  models.DefaultResponse    "Internal Server Error Response"
// @Router       /authors [POST]
func (h *HandlerImpl) CreateAuthor(c *gin.Context) {
	var data models.CreateAuthorModel
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.strg.Author().CreateAuthor(models.CreateAuthorModel{
		Person: data.Person,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Author has been created",
		"data":    nil,
	})
}

func (h *HandlerImpl) UpdateAuthor(c *gin.Context) {
	var data models.Author
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.strg.Author().UpdateAuthor(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author has been updated",
	})
}

func (h *HandlerImpl) DeleteAuthor(c *gin.Context) {
	id := c.Param("id")

	// TODO - delete an author by ID

	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteAuthor",
		"id":      id,
	})
}
