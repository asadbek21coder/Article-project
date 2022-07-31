package postgres

import (
	"fmt"
	"time"

	"project6/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type authorRepoImpl struct {
	db *sqlx.DB
}

func (r authorRepoImpl) CreateAuthor(entity models.CreateAuthorModel) error {
	id := uuid.New()

	createAuthorQuery := `INSERT INTO "author" ("id", "firstname", "lastname") VALUES ($1, $2, $3)`

	result, err := r.db.Exec(createAuthorQuery, id, entity.Firstname, entity.Lastname)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())

	return nil
}

func (r authorRepoImpl) GetAuthorList(queryParams models.QueryParams) (resp models.AuthorList, err error) {
	resp.Authors = []models.Author{}

	params := make(map[string]interface{})
	query := `SELECT
		id,
		firstname,
		lastname,
		created_at,
		updated_at
		FROM author
		`
	filter := " WHERE true"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParams.Search) > 0 {
		params["search"] = queryParams.Search
		filter += " AND ((firstname ILIKE '%' || :search || '%') OR (lastname ILIKE '%' || :search || '%'))"
	}

	if queryParams.Offset > 0 {
		params["offset"] = queryParams.Offset
		offset = " OFFSET :offset"
	}

	if queryParams.Limit > 0 {
		params["limit"] = queryParams.Limit
		limit = " LIMIT :limit"
	}

	cQ := "SELECT count(1) FROM author" + filter
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
		var e models.Author

		err = rows.Scan(
			&e.ID,
			&e.Firstname,
			&e.Lastname,
			&e.CreatedAt,
			&e.UpdateAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Authors = append(resp.Authors, e)
	}

	return resp, nil
}

func (r authorRepoImpl) UpdateAuthor(entity models.UpdateAuthorModel) error {
	query := `
			UPDATE
				author
			SET
				firstname=$1,
				lastname=$2,
				updated_at=$3
			WHERE
				id=$4
	
	`

	_, err := r.db.Exec(query, entity.Firstname, entity.Lastname, time.Now(), entity.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r authorRepoImpl) GetAuthorByID(id string) (resp models.AuthorGetByIDModel, err error) {
	query := `SELECT *
			  FROM author
			  WHERE id=$1	
`
	query2 := `SELECT id,title,body,author_id,created_at,updated_at from article where author_id=$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	for rows.Next() {
		err = rows.Scan(
			&resp.ID,
			&resp.Firstname,
			&resp.Lastname,
			&resp.CreatedAt,
			&resp.UpdateAt,
		)
		if err != nil {
			fmt.Println(err)
			return resp, nil
		}

		rowArticle, err := r.db.Query(query2, id)
		if err != nil {
			fmt.Println(err)
			return resp, err
		}
		for rowArticle.Next() {
			var article models.Article
			err = rowArticle.Scan(
				&article.ID,
				&article.Title,
				&article.Body,
				&article.AuthorID,
				&article.CreatedAt,
				&article.UpdateAt,
			)
			if err != nil {
				fmt.Println(err)
				return resp, nil
			}

			resp.Articles = append(resp.Articles, article)

		}

	}

	return resp, nil
}

func (r authorRepoImpl) DeleteAuthor(id string) error {

	query := `DELETE FROM author WHERE id=$1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
