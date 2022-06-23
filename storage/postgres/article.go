package postgres

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/saidamir98/project6/models"
)

type articleRepoImpl struct {
	db *sqlx.DB
}

var ArticleRepo = articleRepoImpl{}

func (r articleRepoImpl) CloseDB() error {
	return r.db.Close()
}

func init() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=bootcamp password=qwerty123 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	ArticleRepo.db = db
}

func (r articleRepoImpl) CreateArticle(entity models.Article) error {
	checkQuery := `SELECT
		id
	FROM
		"author"
	WHERE
		firstname = $1 AND lastname = $2`

	row, err := r.db.Query(checkQuery, entity.Author.Firstname, entity.Author.Lastname)
	if err != nil {
		return err
	}
	defer row.Close()

	var authorID uuid.UUID
	if row.Next() {
		err = row.Scan(
			&authorID,
		)
		if err != nil {
			return err
		}
	} else {
		authorID = uuid.New()
		createAuthorQuery := `INSERT INTO "author" ("id", "firstname", "lastname") VALUES ($1, $2, $3)`

		result, err := r.db.Exec(createAuthorQuery, authorID, entity.Author.Firstname, entity.Author.Lastname)
		if err != nil {
			return err
		}

		fmt.Println(result.RowsAffected())
	}

	createArticleQuery := `INSERT INTO "article" ("id", "title", "body", "author_id") VALUES ($1, $2, $3, $4)`
	result, err := r.db.Exec(createArticleQuery, uuid.New(), entity.Title, entity.Body, authorID)
	if err != nil {
		return err
	}

	fmt.Println(result.RowsAffected())

	return nil
}

func (r articleRepoImpl) GetArticleList(search string) (resp []models.Article, err error) {
	params := make(map[string]interface{})

	q := `SELECT
	ar.id,
	ar.title,
	ar.body,
	au.firstname,
	au.lastname,
	ar.created_at,
	ar.updated_at
	FROM article AS ar JOIN author au ON au.id = ar.author_id
	`

	if len(search) > 0 {
		params["search"] = search
		q += " WHERE (ar.title ILIKE '%' || :search || '%') OR (ar.body ILIKE '%' || :search || '%')"
	}

	rows, err := r.db.NamedQuery(q, params)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e models.Article

		err = rows.Scan(
			&e.ID,
			&e.Title,
			&e.Body,
			&e.Author.Firstname,
			&e.Author.Lastname,
			&e.CreatedAt,
			&e.UpdateAt,
		)

		// err := rows.StructScan(&e)

		if err != nil {
			return nil, err
		}
		resp = append(resp, e)
	}

	return resp, err
}

func (r articleRepoImpl) UpdateArticle(entity models.Article) error {
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
