package validators

import "regexp"

func IsPasswordLenValid(passwd string) bool {
	return len(passwd) < 8
}

func IsPasswordCharsValid(passwd string) bool {
	const passRegex = ``

	re := regexp.MustCompile(passRegex)
	return re.MatchString(passwd)
}
