package auth

import "errors"

var (
	InvalidBody            = "Invalid body fields!"
	EmailAlreadyRegistered = "Email already registered!"
	FailCreatingUser       = "Failed to create user!"
	InvalidUsername        = "Invalid username!"
	InvalidEmail           = "Invalid e-mail!"
	InvalidPassword        = "Invalid password!"
	InvalidCredentials     = "Invalid credentials!"
	FailInLogin            = "Failed in login!"
	FailGettingUserInfos   = "Failed getting user informations!"
	UserNotLogged          = "User is not logged!"
	TokenNotValid          = "Token JWT is not valid!"
	InvalidPermission      = "User don't have permission to access!"
)

var (
	ErrInvalidBody            = errors.New(InvalidBody)
	ErrEmailAlreadyRegistered = errors.New(EmailAlreadyRegistered)
	ErrFailedCreatingUser     = errors.New(FailCreatingUser)
	ErrInvalidUsername        = errors.New(InvalidUsername)
	ErrInvalidEmail           = errors.New(InvalidEmail)
	ErrInvalidPassword        = errors.New(InvalidPassword)
	ErrInvalidCredentials     = errors.New(InvalidCredentials)
	ErrFailedInLogin          = errors.New(FailInLogin)
	ErrFailedGettingUserInfos = errors.New(FailGettingUserInfos)
	ErrUserNotLogged          = errors.New(UserNotLogged)
	ErrTokenNotValid          = errors.New(TokenNotValid)
	ErrInvalidPermission      = errors.New(InvalidPermission)
)
