package postgres

import (
	"fmt"
	"project6/models"
	"time"

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

	_, err := r.db.Exec(createArticleQuery, id, entity.Title, entity.Body, entity.AuthorID)
	if err != nil {
		return err
	}

	return nil
}

func (r articleRepoImpl) GetArticleList(queryParams models.QueryParams) (resp models.ArticleList, err error) {
	resp.Articles = []models.Article{}

	params := make(map[string]interface{})
	query := `SELECT
		ar.id,
		ar.title,
		ar.body,
		ar.author_id,
		ar.created_at,
		ar.updated_at
		FROM article as ar
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
	query := `SELECT
					ar.id,
					ar.title,
					ar.body,
					at.id as author_id,
					at.firstname,
					at.lastname,
					at.created_at as author_created_at,
					at.updated_at as author_updated_at,		
					ar.created_at,
					ar.updated_at
				FROM
					article as ar
				LEFT JOIN 
					author as at
				ON 
					ar.author_id = at.id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return resp, err
	}
	fmt.Println((rows))
	for rows.Next() {
		var article models.Article
		var author models.Author
		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Body,
			&article.AuthorID,
			&author.Firstname,
			&author.Lastname,
			&author.CreatedAt,
			&author.UpdateAt,
			&article.CreatedAt,
			&article.UpdateAt,
		)

		resp.ID = article.ID
		resp.Title = article.Title
		resp.Body = article.Body
		resp.Author.ID = article.AuthorID
		resp.Author.Firstname = author.Firstname
		resp.Author.Lastname = author.Lastname
		resp.Author.CreatedAt = author.CreatedAt
		resp.Author.UpdateAt = author.UpdateAt
		resp.CreatedAt = article.CreatedAt
		resp.UpdateAt = article.UpdateAt

	}

	return resp, nil
}

func (r articleRepoImpl) UpdateArticle(entity models.UpdateArticleModel) error {
	query := `
	UPDATE
		article
	SET
		title=$1,
		body=$2,
		author_id=$3,
		updated_at=$4
	WHERE
		id=$5
	`
	_, err := r.db.Exec(query, entity.Title, entity.Body, entity.AuthorID, time.Now(), entity.ID)

	if err != nil {
		return err
	}
	return nil
}

func (r articleRepoImpl) DeleteArticle(id string) error {

	query := `DELETE from article where id=$1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}
	return nil
}
