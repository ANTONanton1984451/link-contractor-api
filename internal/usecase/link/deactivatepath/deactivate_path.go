package deactivatepath

import (
	"context"
	"errors"
	"fmt"
)

type (
	DeactivatePath interface {
		Execute(ctx context.Context, path Path, user User) error
	}

	Path struct {
		Path string
	}

	User struct {
		ID int64
	}

	LinkRepo interface {
		UserOwnThisPath(ctx context.Context, path string, userID int64) (bool, error)
		DeactivatePath(ctx context.Context, path string) error
	}
)

var (
	UserIsNotOwnerOfPathErr = errors.New("user dont own this path")
	PathDontExistErr        = errors.New("path dont exist")
)

type usecase struct {
	linkRepo LinkRepo
}

func (uc *usecase) Execute(ctx context.Context, path Path, user User) error {
	isOwn, err := uc.linkRepo.UserOwnThisPath(ctx, path.Path, user.ID)
	if err != nil {
		return fmt.Errorf("chek is owner user: %w", err)
	}
	if !isOwn {
		return UserIsNotOwnerOfPathErr
	}

	err = uc.linkRepo.DeactivatePath(ctx, path.Path)
	if err != nil {
		return fmt.Errorf("deactivate path: %w", err)
	}

	return nil
}
