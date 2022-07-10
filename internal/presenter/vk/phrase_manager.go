package vk

import "fmt"

type phraseManagerPresenter struct{}

func (phraseP *phraseManagerPresenter) DontUnderstand() []byte {
	return []byte(`Не понимаю тебя`)
}

func (phraseP *phraseManagerPresenter) WhatTypeLinks() []byte {
	return []byte(`Какие ссылка мне показывать? Только активные или с неактивынми?`)
}

func (phraseP *phraseManagerPresenter) Greeting(groupLink string) []byte {
	return []byte(fmt.Sprintf(`Привет! Для того, чтобы узнать, кто я и что я могу, можешь зайти в группу %s`, groupLink))
}

func NewPhraseManagerPresenter() *phraseManagerPresenter {
	return &phraseManagerPresenter{}
}
