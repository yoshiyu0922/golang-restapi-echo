package models

import "database/sql"

type UserIdType int

type User struct {
	Id              UserIdType     `json:"id"`
	Name            string         `json:"name"`
	Age             int            `json:"age"`
	JobLargeTypeId  sql.NullString `json:"job_large_type_id"`
	JobMiddleTypeId sql.NullString `json:"job_middle_type_id"`
	JobSmallTypeId  sql.NullString `json:"job_small_type_id"`
	JobName         string         `json:"job_name" boil:"job_name"`
	JobTerm         int            `json:"job_term" boil:"job_term"`
	Messages        Messages
}

type Users []User
