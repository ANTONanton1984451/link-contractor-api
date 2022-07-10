package vk

import (
	"time"
)

type (
	entryPointPresenter struct{}
)

func (mp *entryPointPresenter) InBan(until *time.Time) ([]byte, error) {
	if until != nil {
		return []byte(userInBanUntilTime(*until)), nil
	}
	return []byte(userInPermanentBan()), nil
}

func (mp *entryPointPresenter) UnknownAction() ([]byte, error) {
	return []byte(unknownAction()), nil
}

func NewEntryPointPresenter() *entryPointPresenter {
	return &entryPointPresenter{}
}
