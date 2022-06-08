package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Firstname, Lastname string // Compact by combining the various fields of the same type
}

type Content struct {
	Title string
	Body  string
}

type Article struct {
	ID        int
	Content          // Promoted fields
	Author    Person // Nested structs
	CreatedAt *time.Time
}

var ArticleStorage map[int]Article

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	ArticleStorage = make(map[int]Article)
	now := time.Now()
	ArticleStorage[1] = Article{
		ID: 1,
		Content: Content{
			Title: "title1",
			Body:  "body1",
		},
		Author: Person{
			Firstname: "john1",
			Lastname:  "doe1",
		},
		CreatedAt: &now,
	}

	ArticleStorage[2] = Article{
		ID: 2,
		Content: Content{
			Title: "title2",
			Body:  "body2",
		},
		Author: Person{
			Firstname: "john2",
			Lastname:  "doe2",
		},
		CreatedAt: &now,
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	v1 := api.Group("v1")

	// GetArticleList
	v1.GET("/articles", func(c *gin.Context) {
		var list []Article
		list = make([]Article, 0)

		for _, v := range ArticleStorage {
			list = append(list, v)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "GetArticleList",
			"data":    list,
		})
	})

	// GetArticleByID
	v1.GET("/articles/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "GetArticleByID",
			"id":      id,
		})
	})

	// CreateArticle
	v1.POST("/articles", func(c *gin.Context) {
		var json interface{}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message":   "CreateArticle",
			"json_body": json,
		})
	})

	// UpdateArticle
	v1.PUT("/articles", func(c *gin.Context) {
		var data Article
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, ok := ArticleStorage[data.ID]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
			return
		}

		ArticleStorage[data.ID] = data

		c.JSON(http.StatusOK, gin.H{
			"message":   "CreateArticle",
			"json_body": data,
		})
	})

	// DeleteArticle
	v1.DELETE("/articles/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "DeleteArticle",
			"id":      id,
		})
	})

	r.StaticFS("/static", gin.Dir("static", false))

	r.Run(":7070")
}
