package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}

	// user errors
	ErrUserNotFound      = &Errno{Code: 20101, Message: "The user was not found."}
	ErrPasswordIncorrect = &Errno{Code: 20102, Message: "The password was incorrect."}
)
