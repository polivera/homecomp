package validators

import "fmt"

func IsValidPassword(passwd string) error {
	if len(passwd) < 8 {
		return fmt.Errorf("password needs to be at least 8 character long")
	}
	return nil
}
