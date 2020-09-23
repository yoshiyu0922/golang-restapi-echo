package repository

import (
	"api.com/rest-base-api/src/domain/models"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/dto"
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

func (repo *messageRepository) Search(input *dto.MessageSearchInput) (messages models.Messages, err error) {
	sb := sq.
		Select("id", "title", "message").
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
	err = queries.Raw(query, args...).Bind(context.Background(), repo.sqlHandler.Conn, &messages)
	return
}
