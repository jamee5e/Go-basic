package middlewaresRepositories

import "github.com/jmoiron/sqlx"

type IMiddlewareRepository interface {
}
type middlewaresRepository struct {
	db *sqlx.DB
}

func MiddlewareRepository(db *sqlx.DB) IMiddlewareRepository {
	return &middlewaresRepository{
		db: db,
	}
}
