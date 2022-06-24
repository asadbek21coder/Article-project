package postgres

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/saidamir98/project6/models"
)

type AuthorRepoImpl struct {
	db *sqlx.DB
}

var AuthorRepo = AuthorRepoImpl{}

func (r AuthorRepoImpl) CloseDB() error {
	return r.db.Close()
}

func init() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=bootcamp password=qwerty123 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	AuthorRepo.db = db
}

func (r AuthorRepoImpl) CreateAuthor(entity models.CreateAuthorModel) error {
	id := uuid.New()

	createAuthorQuery := `INSERT INTO "author" ("id", "firstname", "lastname") VALUES ($1, $2, $3)`

	result, err := r.db.Exec(createAuthorQuery, id, entity.Firstname, entity.Lastname)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())

	return nil
}

func (r AuthorRepoImpl) GetAuthorList(queryParams models.QueryParams) (resp models.AuthorList, err error) {
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

func (r AuthorRepoImpl) UpdateAuthor(entity models.Author) error {
	// val, ok := r.db[entity.ID]
	// if !ok {
	// 	return errors.New("not found")
	// }

	// now := time.Now()
	// val.Content = entity.Content
	// val.Author = entity.Author
	// val.UpdateAt = &now

	// r.db[val.ID] = val

	return nil
}
