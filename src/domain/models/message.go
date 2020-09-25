package models

type Message struct {
	Id      int    `json:"message_id"`
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type Messages []Message
