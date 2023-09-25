package appmsg

var (

	// System message
	SUCCESS = "success"
	FAILED  = "failed"

	// System error
	DB_NOT_FOUND          = "db not found"
	BAD_REQUEST           = "bad request"
	INTERNAL_SERVER_ERROR = "internal server error"

	// Business error
	NOT_ENOUGH_MONEY             = "not enough money"
	GET_MAXIMUM_CHANGE_FAILED    = "failed to get maximum change"
	NO_CHANGE                    = "no change"
	NOT_ENOUGH_CHANGE            = "not enough change"
	CALCULATE_CHANGE_FAILED      = "failed to calculate change"
	UPDATE_MAXIMUM_CHANGE_FAILED = "failed to update maximum change"
)
