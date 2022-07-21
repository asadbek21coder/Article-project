package main

import (
	"fmt"

	"project6/config"
	"project6/docs"
	"project6/handlers"
	"project6/storage/postgres"

	"github.com/gin-gonic/gin"

	_ "project6/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath        /api/v1
func main() {

	cfg := config.Load()

	fmt.Printf("%#+v\n", cfg)

	strgPG := postgres.NewPostgres(
		fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresDatabase,
			cfg.PostgresPassword,
		),
	)
	defer strgPG.CloseDB()

	h := handlers.NewHandler(strgPG, cfg)

	switch cfg.Environment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.Title = cfg.ProjectName
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/ping", h.Ping)

	api := r.Group("api")
	v1 := api.Group("v1")
	articleRouter := v1.Group("articles")
	{
		articleRouter.GET("/", h.GetArticleList)
		articleRouter.GET("/:id", h.GetArticleByID)
		articleRouter.POST("/", h.CreateArticle)
		articleRouter.PUT("/", h.UpdateArticle)
		articleRouter.DELETE("/:id", h.DeleteArticle)
	}

	authorRouter := v1.Group("authors")
	{
		authorRouter.GET("/", h.GetAuthorList)
		authorRouter.GET("/:id", h.GetAuthorByID)
		authorRouter.POST("/", h.CreateAuthor)
		authorRouter.PUT("/", h.UpdateAuthor)
		authorRouter.DELETE("/:id", h.DeleteAuthor)
	}

	r.StaticFS("/static", gin.Dir("static", false))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.HTTPPort)
}
