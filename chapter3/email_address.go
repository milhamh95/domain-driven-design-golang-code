package chapter3

import (
	"errors"
	"regexp"
)

type EmailAddress struct {
	value string
}

func NewEmailAddress(value string) (EmailAddress, error) {
	if isValidEmail(value) {
		return EmailAddress{}, errors.New("email is not valid")
	}

	return EmailAddress{
		value: value,
	}, nil
}

func (e EmailAddress) Equal(emailAddress EmailAddress) bool {
	return e == emailAddress
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
