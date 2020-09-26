package repository

import (
	"api.com/go-echo-rest-api/src/domain/models"
	"api.com/go-echo-rest-api/src/infrastructure/database"
	"api.com/go-echo-rest-api/src/interface/dto/input"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/volatiletech/sqlboiler/queries"
)

type UserRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *UserRepository) Search(input *input.UserSearchInput) (users models.Users, err error) {
	sb := sq.Select(`
	u.id,
	name,
	age,
	job_large_type_id,
	job_middle_type_id,
	job_small_type_id,
	job_name,
	job_term
`).From("users u")

	if input.Id != nil {
		sb = sb.Where(sq.Eq{"u.id": input.Id})
	}
	if input.Name != nil {
		sb = sb.Where(sq.Eq{"name": input.Name})
	}
	if input.Age != nil {
		sb = sb.Where(sq.Eq{"age": input.Age})
	}
	if input.MessageId != nil {
		subquery := sq.
			Select("user_id").
			Distinct().
			From("messages").
			Where(sq.Eq{"id": input.MessageId})

		sb = sb.JoinClause(subquery.Prefix("INNER JOIN (").Suffix(") m on m.user_id = u.id"))
	}

	query, args, _ := sb.ToSql()
	err = queries.Raw(query, args...).Bind(context.Background(), repo.SqlHandler.Conn, &users)
	return
}
