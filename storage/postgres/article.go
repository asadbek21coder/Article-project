package postgres

import (
	"fmt"

	"project6/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type articleRepoImpl struct {
	db *sqlx.DB
}

func (r articleRepoImpl) CreateArticle(entity models.CreateArticleModel) error {
	id := uuid.New()

	createArticleQuery := `INSERT INTO "article" ("id", "title", "body", "author_id") VALUES ($1, $2, $3, $4)`

	result, err := r.db.Exec(createArticleQuery, id, entity.Title, entity.Body, entity.AuthorID)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())

	return nil
}

func (r articleRepoImpl) GetArticleList(queryParams models.QueryParams) (resp models.ArticleList, err error) {
	resp.Articles = []models.Article{}

	params := make(map[string]interface{})
	query := `SELECT
		id,
		title,
		body,
		author_id,
		created_at,
		updated_at
		FROM article
		`
	filter := " WHERE true"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParams.Search) > 0 {
		params["search"] = queryParams.Search
		filter += " AND ((title ILIKE '%' || :search || '%') OR (body ILIKE '%' || :search || '%'))"
	}

	if queryParams.Offset > 0 {
		params["offset"] = queryParams.Offset
		offset = " OFFSET :offset"
	}

	if queryParams.Limit > 0 {
		params["limit"] = queryParams.Limit
		limit = " LIMIT :limit"
	}

	cQ := "SELECT count(1) FROM article" + filter
	row, err := r.db.NamedQuery(cQ, params)
	if err != nil {
		return resp, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&resp.Count,
		)
		if err != nil {
			return resp, err
		}
	}

	q := query + filter + offset + limit
	rows, err := r.db.NamedQuery(q, params)
	if err != nil {
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		var e models.Article

		err = rows.Scan(
			&e.ID,
			&e.Title,
			&e.Body,
			&e.AuthorID,
			&e.CreatedAt,
			&e.UpdateAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Articles = append(resp.Articles, e)
	}

	return resp, nil
}

func (r articleRepoImpl) GetArticleByID(id string) (resp models.ArticleFullJoinedModel, err error) {
	// TODO
	return
}

func (r articleRepoImpl) UpdateArticle(entity models.Article) error {
	// TODO
	return nil
}

func (r articleRepoImpl) DeleteArticle(id string) error {
	// TODO
	return nil
}
