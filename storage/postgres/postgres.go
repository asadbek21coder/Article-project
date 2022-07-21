package postgres

import (
	"project6/storage"

	"github.com/jmoiron/sqlx"
)

type postgresImpl struct {
	db          *sqlx.DB
	articleRepo *articleRepoImpl
	authorRepo  *authorRepoImpl
}

func NewPostgres(pgConnStr string) storage.StorageI {
	db, err := sqlx.Connect("postgres", pgConnStr)
	if err != nil {
		panic(err)
	}

	return &postgresImpl{
		db:          db,
		articleRepo: &articleRepoImpl{db: db},
		authorRepo:  &authorRepoImpl{db: db},
	}
}

func (s *postgresImpl) CloseDB() error {
	return s.db.Close()
}

func (s *postgresImpl) Article() storage.ArticleRepoI {
	return s.articleRepo
}

func (s *postgresImpl) Author() storage.AuthorRepoI {
	return s.authorRepo
}
