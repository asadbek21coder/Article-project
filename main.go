package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saidamir98/project6/handlers"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/ping", handlers.Ping)

	api := r.Group("api")
	v1 := api.Group("v1")

	articleRouter := v1.Group("articles")
	{
		articleRouter.GET("/", handlers.GetArticleList)
		articleRouter.GET("/:id", handlers.GetArticleByID)
		articleRouter.POST("/", handlers.CreateArticle)
		articleRouter.PUT("/", handlers.UpdateArticle)
		articleRouter.DELETE("/:id", handlers.DeleteArticle)
	}

	r.StaticFS("/static", gin.Dir("static", false))

	r.Run(":7070")
}
