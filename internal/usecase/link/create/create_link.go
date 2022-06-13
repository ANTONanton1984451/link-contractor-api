package create

import (
	"context"
	"errors"
	"fmt"
	"link-contractor-api/internal/entities/link"
)

// todo валидация пользовательских ссылок
type (
	CreateLink interface {
		Execute(ctx context.Context, link Link, user User) (GeneratedPath, error)
	}

	LinkRepo interface {
		UserHasThisLink(ctx context.Context, linkToCheck string, userID int64) (bool, error)
		InsertLink(ctx context.Context, link NewLink) error
		// Под вопросом, мб такого не будет
		PathExist(ctx context.Context, path string) (bool, error)
	}

	Validation interface {
		ValidLink(link string) (bool, string)
	}

	Link struct {
		Type          LinkType
		Length        int64
		RedirectTo    string
		UserGenerated string
	}

	NewLink struct {
		Type       LinkType
		RedirectTo string
		UserID     int64
		Path       string
	}

	User struct {
		ID int64
	}

	GeneratedPath struct {
		Path string
	}

	LinkType int64
)

var (
	LinkAlreadyExistErr = errors.New("link already exist for this user")
	PathIsBusy          = errors.New("path already busy")
)

const (
	Random LinkType = iota + 1
	UserGenerated
)

type usecase struct {
	linkRepo   LinkRepo
	validation Validation

	createRetryCount int64
}

func New(lr LinkRepo, retryCount int64, validation Validation) CreateLink {
	return &usecase{
		linkRepo:   lr,
		validation: validation,

		createRetryCount: retryCount,
	}
}

func (uc *usecase) Execute(ctx context.Context, link Link, user User) (GeneratedPath, error) {
	valid, rule := uc.validation.ValidLink(link.RedirectTo)
	if !valid {
		return GeneratedPath{}, ValidateErr{
			ValidateRule: rule,
		}
	}
	hasLink, err := uc.linkRepo.UserHasThisLink(ctx, link.RedirectTo, user.ID)
	if err != nil {
		return GeneratedPath{}, fmt.Errorf("check user already has link: %w", err)
	}

	if hasLink {
		return GeneratedPath{}, LinkAlreadyExistErr
	}

	var path string

	switch link.Type {
	case Random:
		randomPath, err := uc.generateRandomPath(ctx, link.Length)
		if err != nil {
			return GeneratedPath{}, fmt.Errorf("generate random path: %w", err)
		}
		path = randomPath
	case UserGenerated:
		exist, err := uc.linkRepo.PathExist(ctx, link.UserGenerated)
		if err != nil {
			return GeneratedPath{}, fmt.Errorf("check exist user path: %w", err)
		}
		if exist {
			return GeneratedPath{}, PathIsBusy
		}

		path = link.UserGenerated
	}

	err = uc.linkRepo.InsertLink(ctx, NewLink{
		Type:       link.Type,
		RedirectTo: link.RedirectTo,
		UserID:     user.ID,
		Path:       path,
	})
	if err != nil {
		return GeneratedPath{}, fmt.Errorf("insert link: %w", err)
	}
	return GeneratedPath{Path: path}, nil
}

func (uc *usecase) generateRandomPath(ctx context.Context, length int64) (string, error) {
	for i := int64(0); i < uc.createRetryCount; i++ {
		path := link.GenerateRandomPath(length)
		// гонка данных
		exist, err := uc.linkRepo.PathExist(ctx, path)
		if err != nil {
			return "", fmt.Errorf("check exist generated path: %w", err)
		}
		if !exist {
			return path, nil
		}
	}

	return "", fmt.Errorf("cannot create uniq path for %d retries", uc.createRetryCount)
}
