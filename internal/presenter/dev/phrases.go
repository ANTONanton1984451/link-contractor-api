package dev

import (
	"fmt"
	"time"
)

func linkAlreadyExistPhrase(link string) string {
	return fmt.Sprintf("Вы уже используете ссылку %s", link)
}

func linkIsBusyPhrase(link string) string {
	return fmt.Sprintf("Ссылка %s уже занята", link)
}

func linkDontExistPhrase(link string) string {
	return fmt.Sprintf("Ссылка %s не существует", link)
}

func userIsNotOwnerLink(link string) string {
	return fmt.Sprintf("Вы не владеете ссылкой %s", link)
}

func validationFailed(rule string) string {
	return fmt.Sprintf("Введённые данные некоректны - %s", rule)
}

func userInPermanentBan() string {
	return fmt.Sprintf("Вы находитесь в вечном бане")
}

func userInBanUntilTime(until time.Time) string {
	return fmt.Sprintf("Вы находитесь в бане до %s", until)
}

func unknownAction() string {
	return "Я не понимаю что ты имеешь в виду"
}

func somethingWentWrong() string {
	return "Что-то пошло не так"
}
