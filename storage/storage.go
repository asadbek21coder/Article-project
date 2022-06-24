package storage

import "github.com/saidamir98/project6/storage/postgres"

type StoreImpl struct {
	Article postgres.ArticleRepoImpl
	Author  postgres.AuthorRepoImpl
}

var Store StoreImpl

func init() {
	Store.Article = postgres.ArticleRepo
	Store.Author = postgres.AuthorRepo
}
