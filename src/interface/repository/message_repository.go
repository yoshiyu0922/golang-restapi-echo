package repository

import (
	"api.com/rest-base-api/src/domain/models"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/usecase"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/volatiletech/sqlboiler/queries"
)

type messageRepository struct {
	sqlHandler *database.SqlHandler
}

func NewMessageRepository(sqlHandler *database.SqlHandler) usecase.MessageRepository {
	return &messageRepository{sqlHandler}
}

func (repo *messageRepository) FindAll() (messages models.Messages, err error) {
	err = queries.Raw("select * from messages").Bind(context.Background(), repo.sqlHandler.Conn, &messages)
	return
}

func (repo *messageRepository) FindByTitle(title *string) (messages models.Messages, err error) {
	sb := sq.
		Select("title", "message").
		From("messages")

	if len(*title) > 0 {
		sb = sb.Where(sq.Eq{"title": *title})
	}

	query, args, _ := sb.ToSql()
	err = queries.Raw(query, args...).Bind(context.Background(), repo.sqlHandler.Conn, &messages)
	return
}
