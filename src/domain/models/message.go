package models

type Message struct {
	Id      int    `json:"message_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type Messages []Message
