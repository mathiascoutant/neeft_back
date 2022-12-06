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

var InvalidInfosProvidedError = Error{
	Code:    "ERR06",
	Message: "Invalid informations provided",
}

var InternalError = Error{
	Code:    "ERR07",
	Message: "Internal Error",
}

var InvalidPriceError = Error{
	Code:    "ERR08",
	Message: "Invalid price",
}

var TournamentNameEmptyError = Error{
	Code:    "ERR09",
	Message: "Tournament name is empty",
}

var AtLeastTwoTeamsError = Error{
	Code:    "ERR10",
	Message: "You must have 2 teams at least",
}

var TournamentWithSameNameUnfinishedError = Error{
	Code:    "ERR11",
	Message: "A tournament with the same name already exists and isn't finished",
}

var InvalidTeamSizeError = Error{
	Code:    "ERR12",
	Message: "The team size is invalid",
}

var InvalidPartyModeError = Error{
	Code:    "ERR13",
	Message: "Invalid party mode (Fortnite)",
}

var PasswordTooShortError = Error{
	Code:    "ERR14",
	Message: "Password is too short",
}

var InvalidFirstNameError = Error{
	Code:    "ERR15",
	Message: "Invalid first name",
}

var InvalidLastNameError = Error{
	Code:    "ERR16",
	Message: "Invalid last name",
}

var InvalidEmailError = Error{
	Code:    "ERR17",
	Message: "Invalid email address",
}

var AccountAlreadyExistError = Error{
	Code:    "ERR18",
	Message: "An account with the same username or email has already been created",
}

var InvalidDateTimeError = Error{
	Code:    "ERR19",
	Message: "Invalid date or time value",
}

var TournamentDoesNotExistError = Error{
	Code:    "ERR20",
	Message: "Provided tournament does not exist",
}
