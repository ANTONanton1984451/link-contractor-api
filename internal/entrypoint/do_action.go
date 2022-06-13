package entrypoint

import (
	"context"
	"errors"
	"fmt"
)

func (m *Entrypoint) DoAction(ctx context.Context, userExternalID int64, input []byte) ([]byte, error) {
	user, err := m.auth.GetUser(ctx, userExternalID)
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

	resp, err := action(user)
	if err != nil {
		return nil, fmt.Errorf("execute action: %w", err)
	}

	return resp.Output, nil
}
