package redirect

import (
	"context"
	"errors"
)

type redirect struct{}

// todo implement this
func (rd *redirect) Store(ctx context.Context, path string, redirectTo string) error {
	return errors.New("test error")
}

func New() *redirect {
	return &redirect{}
}
