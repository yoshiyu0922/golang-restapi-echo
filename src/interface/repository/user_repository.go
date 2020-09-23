package repository

import (
	"api.com/rest-base-api/src/domain/models"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/dto"
	"api.com/rest-base-api/src/usecase"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/volatiletech/sqlboiler/queries"
)

type userRepository struct {
	sqlHandler *database.SqlHandler
}

func NewUserRepository(sqlHandler *database.SqlHandler) usecase.UserRepository {
	return &userRepository{sqlHandler}
}

func (repo *userRepository) Search(input *dto.UserSearchInput) (users models.Users, err error) {
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
	fmt.Println(query)
	err = queries.Raw(query, args...).Bind(context.Background(), repo.sqlHandler.Conn, &users)
	return
}
