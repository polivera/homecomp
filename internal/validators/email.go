package validators

import (
	"fmt"
	"regexp"
)

func EmailValidator(email string) (bool, []error) {
	errorList := make([]error, 2)

	if !emailStringValidator(email) {
		errorList = append(errorList, fmt.Errorf("invalid email address"))
	}

	return len(errorList) == 0, errorList
}

func emailStringValidator(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}
