package helper

import (
	"strings"
)

var (
	// MinBirthDateYear ...
	MinBirthDateYear = 18
)

// PhoneValidateAndConcat ...
func PhoneValidateAndConcat(code, phone string) (res string) {
	phone = strings.Replace(phone, code, "", -1)
	res = code + phone

	return res
}

// EmailReplaceCountryCode ...
func EmailReplaceCountryCode(email, countryCode string) (res string) {
	if strings.Contains(email, countryCode) {
		res = strings.ReplaceAll(email, countryCode, "")
	} else {
		res = email
	}

	return res
}
