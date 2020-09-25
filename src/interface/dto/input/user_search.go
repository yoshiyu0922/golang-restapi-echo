package input

type UserSearchInput struct {
	Id              *int    `query:"user_id"`
	Name            *string `query:"name"`
	Age             *int    `query:"age"`
	JobLargeTypeId  *string `query:"job_large_type_id"`
	JobMiddleTypeId *string `query:"job_middle_type_id"`
	JobSmallTypeId  *string `query:"job_small_type_id"`
	JobName         *string `query:"job_name"`
	JobTerm         *int    `query:"job_term"`
	MessageId       *int    `query:"message_id"`
}
