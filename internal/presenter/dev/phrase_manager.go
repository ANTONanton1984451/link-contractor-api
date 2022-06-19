package dev

import "encoding/json"

type phraseManagerPresenter struct{}

func (phraseP *phraseManagerPresenter) DontUnderstand() []byte {
	resp, _ := json.Marshal(OkResponse{Status: _okStatus, Result: `Я не понимаю тебя`})
	return resp
}

func (phraseP *phraseManagerPresenter) WhatTypeLinks() []byte {
	resp, _ := json.Marshal(OkResponse{Status: _okStatus, Result: `Какие ссылка мне показывать? Только активные или с неактивынми?`})
	return resp
}

func NewPhraseManagerPresenter() *phraseManagerPresenter {
	return &phraseManagerPresenter{}
}
