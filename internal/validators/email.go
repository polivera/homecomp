package validators

import (
	"fmt"
	"regexp"
)

func IsEmailValid(email string) error {
	if !emailStringValidator(email) {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

func emailStringValidator(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}
