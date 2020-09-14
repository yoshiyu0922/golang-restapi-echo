package models

type Message struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type Messages []Message
