package dto

type MessageSearchInput struct {
	Id      *string `query:"message_id"`
	UserId  *string `query:"user_id"`
	Title   *string `query:"title"`
	Message *string `query:"message"`
}
