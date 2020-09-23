package models

type User struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Age             int    `json:"age"`
	JobLargeTypeId  string `json:"job_large_type_id"`
	JobMiddleTypeId string `json:"job_middle_type_id"`
	JobSmallTypeId  string `json:"job_small_type_id"`
	JobName         string `json:"job_name"`
	JobTerm         int    `json:"job_term"`
	Messages        Messages
}

type Users = []User
