package storage

import (
	"github.com/saidamir98/project6/models"
	"github.com/saidamir98/project6/storage/postgres"
)

var Store = postgres.ArticleRepo

func init() {
	err := Store.CreateArticle(models.Article{
		Content: models.Content{
			Title: "title1",
			Body:  "body1",
		},
		Author: models.Person{
			Firstname: "john1",
			Lastname:  "doe1",
		},
	})

	if err != nil {
		panic(err)
	}

	err = Store.CreateArticle(models.Article{
		Content: models.Content{
			Title: "title2",
			Body:  "body2",
		},
		Author: models.Person{
			Firstname: "john2",
			Lastname:  "doe2",
		},
	})
	if err != nil {
		panic(err)
	}
}
