package dev

import "encoding/json"

type handlerPresenter struct{}

func (hp *handlerPresenter) SomethingWentWrong() []byte {
	resp, _ := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: somethingWentWrong(),
	})

	return resp
}

func NewHandlerPresenter() *handlerPresenter {
	return &handlerPresenter{}
}
