package register

import (
	"context"
	"errors"
)

// todo не актуально
type (
	RegisterUser interface {
		Execute(context.Context, User) error
	}

	UserRepo interface {
		CreateUser(context.Context, User) error
	}

	User struct {
		ID         int64
		ExternalID int64
		Name       string
		SurName    string
		SourceID   int64
	}
)

var (
	UserAlreadyExistErr = errors.New("user already exist")
)

type usecase struct {
	userRepo UserRepo
}

func (uc *usecase) Execute(ctx context.Context, guest User) error {
	if guest.ExternalID != 0 {
		return UserAlreadyExistErr
	}

	return uc.userRepo.CreateUser(ctx, guest)
}
