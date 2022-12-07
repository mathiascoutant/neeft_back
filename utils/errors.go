package utils

var InvalidRequestFormat = Error{
	StatusCode: 400,
	ErrCode:    "ERR01",
	Message:    "Invalid request format",
}

var UsernameEmptyError = Error{
	StatusCode: 400,
	ErrCode:    "ERR02",
	Message:    "Username is empty",
}

var PasswordEmptyError = Error{
	StatusCode: 400,
	ErrCode:    "ERR03",
	Message:    "Password is empty",
}

var UsernameOrPasswordInvalidError = Error{
	StatusCode: 400,
	ErrCode:    "ERR04",
	Message:    "Username or password is invalid",
}

var TeamWithSameNameExistsError = Error{
	StatusCode: 400,
	ErrCode:    "ERR05",
	Message:    "A team with the same name already exist",
}

var InvalidInfosProvidedError = Error{
	StatusCode: 400,
	ErrCode:    "ERR06",
	Message:    "Invalid informations provided",
}

var InternalError = Error{
	StatusCode: 400,
	ErrCode:    "ERR07",
	Message:    "Internal Error",
}

var InvalidPriceError = Error{
	StatusCode: 400,
	ErrCode:    "ERR08",
	Message:    "Invalid price",
}

var TournamentNameEmptyError = Error{
	StatusCode: 400,
	ErrCode:    "ERR09",
	Message:    "Tournament name is empty",
}

var AtLeastTwoTeamsError = Error{
	StatusCode: 400,
	ErrCode:    "ERR10",
	Message:    "You must have 2 teams at least",
}

var TournamentWithSameNameUnfinishedError = Error{
	StatusCode: 400,
	ErrCode:    "ERR11",
	Message:    "A tournament with the same name already exists and isn't finished",
}

var InvalidTeamSizeError = Error{
	StatusCode: 400,
	ErrCode:    "ERR12",
	Message:    "The team size is invalid",
}

var InvalidPartyModeError = Error{
	StatusCode: 400,
	ErrCode:    "ERR13",
	Message:    "Invalid party mode (Fortnite)",
}

var PasswordTooShortError = Error{
	StatusCode: 400,
	ErrCode:    "ERR14",
	Message:    "Password is too short",
}

var InvalidFirstNameError = Error{
	StatusCode: 400,
	ErrCode:    "ERR15",
	Message:    "Invalid first name",
}

var InvalidLastNameError = Error{
	StatusCode: 400,
	ErrCode:    "ERR16",
	Message:    "Invalid last name",
}

var InvalidEmailError = Error{
	StatusCode: 400,
	ErrCode:    "ERR17",
	Message:    "Invalid email address",
}

var AccountAlreadyExistError = Error{
	StatusCode: 400,
	ErrCode:    "ERR18",
	Message:    "An account with the same username or email has already been created",
}

var InvalidDateTimeError = Error{
	StatusCode: 400,
	ErrCode:    "ERR19",
	Message:    "Invalid date or time value",
}

var TournamentDoesNotExistError = Error{
	StatusCode: 400,
	ErrCode:    "ERR20",
	Message:    "Provided tournament does not exist",
}
