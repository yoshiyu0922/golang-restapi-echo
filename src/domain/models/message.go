package models

type MessageIdType int

type Message struct {
	Id      MessageIdType `json:"message_id"`
	UserId  UserIdType    `json:"user_id"`
	Title   string        `json:"title"`
	Message string        `json:"message"`
}

type Messages []Message
