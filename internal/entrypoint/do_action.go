package entrypoint

import (
	"context"
	"errors"
	"fmt"
	"link-contractor-api/internal/entities/user"
)

// DoAction вход в бизнес-логику приложения,
// Entrypoint получает пользователя, получает у менеджера фраз нужный экш и отдаёт результат выполнения экшена
func (m *Entrypoint) DoAction(ctx context.Context, usr user.User, input []byte) ([]byte, error) {
	usr, err := m.auth.GetUser(ctx, usr)
	if err != nil {
		var uBan UserInBanErr
		if errors.As(err, &uBan) {
			return m.presenter.InBan(uBan.Until)
		}

		return nil, fmt.Errorf("get user: %w", err)
	}

	action, err := m.phManager.GetAction(ctx, input)
	if err != nil {
		if errors.Is(UnknownActionErr, err) {
			return m.presenter.UnknownAction()
		}
		return nil, fmt.Errorf("get action: %w", err)
	}

	resp, err := action(usr)
	if err != nil {
		return nil, fmt.Errorf("execute action: %w", err)
	}

	return resp.Output, nil
}
