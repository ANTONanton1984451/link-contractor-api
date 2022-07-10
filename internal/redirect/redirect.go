package redirect

import (
	"context"
)

type redirect struct{}

// todo implement this
func (rd *redirect) Store(ctx context.Context, path string, redirectTo string) error {
	return nil
}

func New() *redirect {
	return &redirect{}
}
