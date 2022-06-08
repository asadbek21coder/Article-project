package storage

import (
	"github.com/saidamir98/project6/models"
	"github.com/saidamir98/project6/storage/inmemory"
)

var Store = inmemory.ArticleRepo

func init() {
	Store.CreateArticle(models.Article{
		ID: "1",
		Content: models.Content{
			Title: "title1",
			Body:  "body1",
		},
		Author: models.Person{
			Firstname: "john1",
			Lastname:  "doe1",
		},
	})

	Store.CreateArticle(models.Article{
		ID: "2",
		Content: models.Content{
			Title: "title2",
			Body:  "body2",
		},
		Author: models.Person{
			Firstname: "john2",
			Lastname:  "doe2",
		},
	})
}
