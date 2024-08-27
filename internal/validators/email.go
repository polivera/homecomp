package validators

import (
	"context"
	"regexp"

	"homecomp/internal/repositories"
)

func IsEmailStringValid(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}

func IsEmailNew(ctx context.Context, email string, repo repositories.UserRepo) bool {
	user := repo.GetUserByEmail(ctx, email)
	return user.ID == 0
}
