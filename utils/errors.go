package utils

var InvalidRequestFormat = Error{
	Code:    "ERR01",
	Message: "Invalid request format",
}

var UsernameEmptyError = Error{
	Code:    "ERR02",
	Message: "Username is empty",
}

var PasswordEmptyError = Error{
	Code:    "ERR03",
	Message: "Password is empty",
}

var UsernameOrPasswordInvalidError = Error{
	Code:    "ERR04",
	Message: "Username or password is invalid",
}

var TeamWithSameNameExistsError = Error{
	Code:    "ERR05",
	Message: "A team with the same name already exist",
}

var InvalidInfosProvided = Error{
	Code:    "ERR06",
	Message: "Invalid informations provided",
}

var DatabaseError = Error{
	Code:    "ERR07",
	Message: "Database Error",
}
