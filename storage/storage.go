package storage

import "project6/models"

type StorageI interface {
	Article() ArticleRepoI
	Author() AuthorRepoI
	CloseDB() error
}

type ArticleRepoI interface {
	CreateArticle(entity models.CreateArticleModel) error
	GetArticleList(queryParams models.QueryParams) (resp models.ArticleList, err error)
	GetArticleByID(id string) (resp models.ArticleFullJoinedModel, err error)
	UpdateArticle(entity models.UpdateArticleModel) error
	DeleteArticle(id string) error
}

type AuthorRepoI interface {
	CreateAuthor(entity models.CreateAuthorModel) error
	GetAuthorList(queryParams models.QueryParams) (resp models.AuthorList, err error)
	// GetAuthorByID(id string) (resp models.AuthorGetByIDModel, err error)
	UpdateAuthor(entity models.Author) error
	// DeleteAuthor(id string) error
}

// type StoreImpl struct {
// 	Article postgres.ArticleRepoImpl
// 	Author  postgres.AuthorRepoImpl
// }

// var Store StoreImpl

// func init() {
// 	Store.Article = postgres.ArticleRepo
// 	Store.Author = postgres.AuthorRepo
// }
