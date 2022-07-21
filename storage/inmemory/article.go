package inmemory

import (
	"errors"
	"time"

	"project6/models"
)

type articleRepoImpl struct {
	db map[string]models.Article
}

var ArticleRepo = articleRepoImpl{}

func init() {
	ArticleRepo.db = make(map[string]models.Article)
}

func (r articleRepoImpl) CreateArticle(entity models.Article) error {
	_, ok := r.db[entity.ID]
	if ok {
		return errors.New("already exists")
	}

	now := time.Now()
	entity.CreatedAt = &now

	r.db[entity.ID] = entity

	return nil
}

func (r articleRepoImpl) GetArticleList(search string) (resp []models.Article) {

	for _, v := range r.db {
		// TODO - filter result based on 'search' query param
		resp = append(resp, v)
	}

	return resp
}

func (r articleRepoImpl) UpdateArticle(entity models.Article) error {
	val, ok := r.db[entity.ID]
	if !ok {
		return errors.New("not found")
	}

	now := time.Now()
	val.Content = entity.Content
	val.AuthorID = entity.AuthorID
	val.UpdateAt = &now

	r.db[val.ID] = val

	return nil
}
