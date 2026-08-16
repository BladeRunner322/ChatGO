package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/BladeRunner322/ChatGO/internal/core/errors"
)

type User struct {
	ID           int
	Version      int
	NickName     string
	PasswordHash string
}

func NewUser(
	id int,
	version int,
	nickName string,
	passwordHash string,
) User {
	return User{
		ID:           id,
		Version:      version,
		NickName:     nickName,
		PasswordHash: passwordHash,
	}
}

func NewUserUninitialised(
	nickName string,
	passwordHash string,
) User {
	return User{
		ID:           UninitializedID,
		Version:      UninitializedVersion,
		NickName:     nickName,
		PasswordHash: passwordHash,
	}
}

func (u *User) Validate() error {
	nickNameLen := len([]rune(u.NickName))
	if nickNameLen < 3 || nickNameLen > 15 {
		return fmt.Errorf(
			"invalid 'NickName' len: %d: %w",
			nickNameLen,
			core_errors.ErrInvalidArgument,
		)
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	if !re.MatchString(u.NickName) {
		return fmt.Errorf(
			"invalid 'NickName' format: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
