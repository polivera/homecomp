package validators

import (
	"regexp"

	"homecomp/internal/repositories"
)

func IsEmailStringValid(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}

func IsEmailNew(email string, repo repositories.UserRepo) bool {
	user := repo.GetUserByEmail(email)
	return user.ID == 0
}
