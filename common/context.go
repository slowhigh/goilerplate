package common

const (
	ACCESS_TOKEN     = "access_token"
	ACCESS_TOKEN_TTL = 60 * 30 // sec

	REFRESH_TOKEN     = "refresh_token"
	REFRESH_TOKEN_TTL = 60 * 60 * 24 //sec

	CURRENT_USER_ID = "curUserID"

	ERR_EMPTY_USER = "the current user id is empty"
	ERR_EMPTY_PARAM = "the param is empty"
)
