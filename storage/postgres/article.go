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

var dbSchema = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "article" (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
	"title" varchar,
	"body" text,
	"author_id" UUID,
	"created_at" timestamp NOT NULL DEFAULT NOW(),
	"updated_at" timestamp NOT NULL DEFAULT NOW()
  );

  CREATE UNIQUE INDEX unique_title_on_article ON "article" ("title");
  
  CREATE TABLE IF NOT EXISTS "author" (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
	"firstname" varchar,
	"lastname" varchar,
	"created_at" timestamp NOT NULL DEFAULT NOW(),
	"updated_at" timestamp NOT NULL DEFAULT NOW()
  );
  
  CREATE UNIQUE INDEX unique_firstname_lastname_on_author ON "author" ("firstname", "lastname");
  
  ALTER TABLE "article" ADD CONSTRAINT article_author_id_fkey FOREIGN KEY ("author_id") REFERENCES "author" ("id") ON DELETE CASCADE;
`

func (r articleRepoImpl) CloseDB() error {
	return r.db.Close()
}

func init() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=bootcamp password=qwerty123 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(dbSchema)
	db.MustExec(`DELETE FROM "author"`)
	db.MustExec(`DELETE FROM "article"`)

	tx := db.MustBegin()

	tx.MustExec(`INSERT INTO "author" ("firstname", "lastname") VALUES ($1, $2)`, "John", "Doe")

	tx.MustExec(`INSERT INTO "author" ("firstname", "lastname") VALUES ($1, $2)`, "Steve", "Jobs")

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		panic(err)
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

func (r articleRepoImpl) GetArticleList(search string) (resp []models.Article) {
	//select ar.id, ar.title, ar.body, au.firstname, au.lastname from article as ar join author au on au.id = ar.author_id;
	// for _, v := range r.db {
	// 	// TODO - filter result based on 'search' query param
	// 	resp = append(resp, v)
	// }

	return resp
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
