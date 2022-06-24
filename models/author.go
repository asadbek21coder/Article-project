package models

import "time"

type Author struct {
	ID string `json:"id"`
	Person
	CreatedAt *time.Time `json:"created_at"`
	UpdateAt  *time.Time `json:"update_at"`
}

type CreateAuthorModel struct {
	Person
}

type UpdateAuthorModel struct {
	ID string `json:"id"`
	Person
}

type AuthorList struct {
	Authors []Author `json:"authors"`
	Count   int      `json:"count"`
}
