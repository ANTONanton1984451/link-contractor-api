package entrypoint

import (
	"context"
	"errors"
	"fmt"
	"link-contractor-api/internal/controllers"
	"time"
)

// todo переименовать

type (
	PhraseManager interface {
		GetAction(ctx context.Context, input []byte) (ActionFunc, error)
	}

	Auth interface {
		GetUser(ctx context.Context, externalID int64) (User, error)
	}

	Presenter interface {
		InBan(until *time.Time) ([]byte, error)
		UnknownAction() ([]byte, error)
	}

	// todo мб  унести в ентитю
	User struct {
		Name         string
		SurName      string
		ID           int64
		ExternalID   int64
		RegisteredAt time.Time
	}

	ActionFunc func(user User) (controllers.Response, error)

	Entrypoint struct {
		auth      Auth
		phManager PhraseManager
		presenter Presenter
	}

	UserInBanErr struct {
		Until *time.Time
	}
)

func New(auth Auth, phManager PhraseManager, presenter Presenter) Entrypoint {
	return Entrypoint{
		auth:      auth,
		phManager: phManager,
		presenter: presenter,
	}
}

var (
	UnknownActionErr = errors.New("unknown action")
)

func (uBan UserInBanErr) Error() string {
	if uBan.Until != nil {
		return fmt.Sprintf("user in ban until %s", uBan.Until.String())
	}

	return fmt.Sprintf("user in permanent ban")
}
