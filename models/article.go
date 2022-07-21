package models

import "time"

type Content struct {
	Title string `json:"title" default:"Lorem"`
	Body  string `json:"body"`
}

type Article struct {
	ID string `json:"id"`
	Content
	AuthorID  string     `json:"author_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdateAt  *time.Time `json:"update_at"`
}

type ArticleFullJoinedModel struct {
	ID string `json:"id"`
	Content
	Author    Author     `json:"author"`
	CreatedAt *time.Time `json:"created_at"`
	UpdateAt  *time.Time `json:"update_at"`
}

type CreateArticleModel struct {
	Content
	AuthorID string `json:"author_id"`
}

type UpdateArticleModel struct {
	ID string `json:"id"`
	Content
	AuthorID string `json:"author_id"`
}

type ArticleList struct {
	Articles []Article `json:"articles"`
	Count    int       `json:"count"`
}
