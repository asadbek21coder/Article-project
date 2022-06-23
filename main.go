package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saidamir98/project6/handlers"
	"github.com/saidamir98/project6/storage/postgres"

	_ "github.com/saidamir98/project6/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.1
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:7070
// @BasePath        /api/v1
func main() {
	defer postgres.ArticleRepo.CloseDB()

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":7070")
}
