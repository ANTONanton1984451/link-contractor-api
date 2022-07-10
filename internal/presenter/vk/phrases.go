package vk

import (
	"fmt"
	"link-contractor-api/internal/controllers"
	"time"
)

func linkAlreadyExistPhrase(link string) string {
	return fmt.Sprintf("Вы уже используете ссылку %s", link)
}

func youDontHaveLinks() string {
	return `На данный момент у тебя нет ссылок`
}

func formLink(lnk controllers.ListLink) string {
	message := fmt.Sprintf("Ссылка %s, ведеёт на %s, создана %s", lnk.From, lnk.To, lnk.CreatedAt)
	if !lnk.Active {
		message += " , не активна"
	}

	return message
}

func yourLinkPhrase(link string) string {
	return fmt.Sprintf("Ваша ссылка - %s", link)
}

func linkIsBusyPhrase(link string) string {
	return fmt.Sprintf("Ссылка %s уже занята", link)
}

func linkDontExistPhrase(link string) string {
	return fmt.Sprintf("Ссылка %s не существует", link)
}

func userIsNotOwnerLinkPhrase(link string) string {
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
	return "Извини, что-то пошло не так"
}
