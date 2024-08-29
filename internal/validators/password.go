package validators

import "regexp"

func IsPasswordLenValid(passwd string) bool {
	return len(passwd) < 8
}

func IsPasswordCharsValid(passwd string) bool {
	return regexp.MustCompile(`[^a-zA-Z\d]`).MatchString(passwd) &&
		regexp.MustCompile(`[A-Z]`).MatchString(passwd) &&
		regexp.MustCompile(`\d`).MatchString(passwd) &&
		regexp.MustCompile(`[a-z]`).MatchString(passwd)
}
