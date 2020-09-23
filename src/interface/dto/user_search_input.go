package dto

type UserSearchInput struct {
	Id              *int    `json:"user_id" query:"user_id"`
	Name            *string `json:"name" query:"name"`
	Age             *int    `json:"age" query:"age"`
	JobLargeTypeId  *string `json:"job_large_type_id" query:"job_large_type_id"`
	JobMiddleTypeId *string `json:"job_middle_type_id" query:"job_middle_type_id"`
	JobSmallTypeId  *string `json:"job_small_type_id" query:"job_small_type_id"`
	JobName         *string `json:"job_name" query:"job_name"`
	JobTerm         *int    `json:"job_term" query:"job_term"`
	MessageId       *int    `json:"message_id" query:"message_id"`
}
