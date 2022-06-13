package dev

import (
	"encoding/json"
	"time"
)

type (
	entryPointPresenter struct{}
)

func (mp *entryPointPresenter) InBan(until *time.Time) ([]byte, error) {
	response := FailResponse{
		Status: _failStatus,
	}

	if until != nil {
		response.Reason = userInBanUntilTime(*until)
	} else {
		response.Reason = userInPermanentBan()
	}

	return json.Marshal(response)
}

func (mp *entryPointPresenter) UnknownAction() ([]byte, error) {
	return json.Marshal(FailResponse{Status: _failStatus, Reason: unknownAction()})
}

func NewEntryPointPresenter() *entryPointPresenter {
	return &entryPointPresenter{}
}
