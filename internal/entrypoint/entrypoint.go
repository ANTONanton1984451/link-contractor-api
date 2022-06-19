package entrypoint

import (
	"context"
	"errors"
	"fmt"
	"link-contractor-api/internal/controllers"
	"link-contractor-api/internal/entities/user"
	"time"
)

type (
	PhraseManager interface {
		GetAction(ctx context.Context, input []byte) (ActionFunc, error)
	}

	Auth interface {
		GetUser(ctx context.Context, usr user.User) (user.User, error)
	}

	Presenter interface {
		InBan(until *time.Time) ([]byte, error)
		UnknownAction() ([]byte, error)
	}

	ActionFunc func(user user.User) (controllers.Response, error)

	Entrypoint struct {
		auth      Auth
		phManager PhraseManager
		presenter Presenter
	}

	SourceType uint64

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
