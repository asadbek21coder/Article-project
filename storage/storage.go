package storage

import (
	"fmt"

	"github.com/saidamir98/project6/storage/postgres"
)

var Store = postgres.ArticleRepo

func init() {
	fmt.Println("Storage init")
}
