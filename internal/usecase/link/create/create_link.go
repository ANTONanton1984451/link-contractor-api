package create

import (
	"context"
	"errors"
	"fmt"
	"link-contractor-api/internal/entities/link"
)

type (
	CreateLink interface {
		Execute(ctx context.Context, link Link, user User) (GeneratedPath, error)
	}

	LinkRepo interface {
		UserHasThisLink(ctx context.Context, linkToCheck string, userID int64) (bool, error)
		InsertLink(ctx context.Context, link NewLink) error
	}

	Validation interface {
		ValidLink(link string) (bool, string)
		ValidPath(path string) (bool, string)
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

// Execute сценарий создания ссылки, сначала проверяем что пользователь уже имеет данную ссылку, а азтем генерируем сылку на основе запроса -
// либо генерим рандомную ссылку, либо создаёт ссылку, которую указал пользователь
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
	newLink := NewLink{
		Type:       link.Type,
		RedirectTo: link.RedirectTo,
		UserID:     user.ID,
	}

	switch link.Type {
	case Random:
		path, err = uc.insetRandomPath(ctx, link.Length, newLink)
		if err != nil {
			return GeneratedPath{}, fmt.Errorf("insert link with random path: %w", err)
		}
	case UserGenerated:
		valid, rule = uc.validation.ValidPath(link.UserGenerated)
		if !valid {
			return GeneratedPath{}, ValidateErr{
				ValidateRule: rule,
			}
		}

		newLink.Path = link.UserGenerated

		err = uc.linkRepo.InsertLink(ctx, newLink)
		if err != nil {
			return GeneratedPath{}, fmt.Errorf("insert link with user path: %w", err)
		}
		path = link.UserGenerated
	}

	return GeneratedPath{Path: path}, nil
}

func (uc *usecase) insetRandomPath(ctx context.Context, length int64, newLink NewLink) (string, error) {
	for i := int64(0); i < uc.createRetryCount; i++ {
		path := link.GenerateRandomPath(length)

		newLink.Path = path
		err := uc.linkRepo.InsertLink(ctx, newLink)
		if err != nil {
			if errors.Is(PathIsBusy, err) {
				continue
			}
			return "", fmt.Errorf("insert link: %w", err)
		}

		return path, nil
	}

	return "", fmt.Errorf("cannot create uniq path for %d retries", uc.createRetryCount)
}
