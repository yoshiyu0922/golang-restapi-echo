package repository

import (
	"api.com/go-echo-rest-api/src/adapter/dto/input"
	"api.com/go-echo-rest-api/src/domain/models"
	"api.com/go-echo-rest-api/src/infrastructure/database"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/volatiletech/sqlboiler/queries"
)

type MessageRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *MessageRepository) FindAll() (messages models.Messages, err error) {
	err = queries.Raw("select * from messages").Bind(context.Background(), repo.SqlHandler.Conn, &messages)
	return
}

func (repo *MessageRepository) Search(input *input.MessageSearchInput) (messages models.Messages, err error) {
	sb := sq.
		Select("id", "user_id", "title", "message").
		From("messages")

	if input.Id != nil {
		sb = sb.Where(sq.Eq{"id": input.Id})
	}
	if input.UserId != nil {
		sb = sb.Where(sq.Eq{"user_id": input.UserId})
	}
	if input.Title != nil {
		sb = sb.Where(sq.Eq{"title": input.Title})
	}
	if input.Message != nil {
		sb = sb.Where(sq.Eq{"message": input.Message})
	}

	query, args, _ := sb.ToSql()
	err = queries.Raw(query, args...).Bind(context.Background(), repo.SqlHandler.Conn, &messages)
	return
}

func (repo *MessageRepository) FindByUserIds(userIds *[]models.UserIdType) (messages models.Messages, err error) {
	sb := sq.
		Select("id", "user_id", "title", "message").
		From("messages").
		Where(sq.Eq{"user_id": userIds})

	query, args, _ := sb.ToSql()
	err = queries.Raw(query, args...).Bind(context.Background(), repo.SqlHandler.Conn, &messages)
	return
}
