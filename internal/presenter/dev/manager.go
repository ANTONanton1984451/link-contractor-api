package dev

import (
	"encoding/json"
	"time"
)

type (
	entryPointPresenter struct{}
)

func (mp *entryPointPresenter) InBan(until *time.Time) ([]byte, error) {
	resp := FailResponse{
		Status: _failStatus,
	}

	if until != nil {
		resp.Reason = userInBanUntilTime(*until)
	} else {
		resp.Reason = userInPermanentBan()
	}

	return json.Marshal(resp)
}

func (mp *entryPointPresenter) UnknownAction() ([]byte, error) {
	return json.Marshal(FailResponse{Status: _failStatus, Reason: unknownAction()})
}

func NewEntryPointPresenter() *entryPointPresenter {
	return &entryPointPresenter{}
}
