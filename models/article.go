package models

import "time"

type Content struct {
	Title string `json:"title" default:"Lorem"`
	Body  string `json:"body"`
}

type Article struct {
	ID string `json:"id"`
	Content
	Author    Person     `json:"author"`
	CreatedAt *time.Time `json:"created_at"`
	UpdateAt  *time.Time `json:"update_at"`
}

type CreateArticleModel struct {
	Content
	Author Person `json:"author"`
}

type UpdateArticleModel struct {
	ID string `json:"id"`
	Content
	Author Person `json:"author"`
}
