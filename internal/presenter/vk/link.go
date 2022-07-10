package vk

import (
	"link-contractor-api/internal/controllers"
	"link-contractor-api/internal/response"
	"strings"
)

type (
	linkPresenter struct{}
)

func (p *linkPresenter) ShowGeneratedLink(link controllers.GeneratedLink) (response.DTO, error) {
	return response.DTO{Output: []byte(yourLinkPhrase(link.Link))}, nil
}

func (p *linkPresenter) ActivateOK() (response.DTO, error) {
	return response.DTO{Output: []byte(`Ссылка активирована`)}, nil
}

func (p *linkPresenter) DeactivateOK() (response.DTO, error) {
	return response.DTO{Output: []byte(`Ссылка деактивирована`)}, nil
}

func (p *linkPresenter) ListLinks(linkList []controllers.ListLink) (response.DTO, error) {
	if len(linkList) == 0 {
		return response.DTO{Output: []byte(youDontHaveLinks())}, nil
	}

	b := strings.Builder{}
	for _, lnk := range linkList {
		b.WriteString("• ")
		b.WriteString(formLink(lnk))
		b.WriteString("\n")
	}

	return response.DTO{Output: []byte(b.String())}, nil
}

func (p *linkPresenter) LinkAlreadyExist(link string) (response.DTO, error) {
	return response.DTO{Output: []byte(linkAlreadyExistPhrase(link))}, nil
}

func (p *linkPresenter) LinkIsBusy(link string) (response.DTO, error) {
	return response.DTO{Output: []byte(linkIsBusyPhrase(link))}, nil
}

func (p *linkPresenter) LinkDontExist(link string) (response.DTO, error) {
	return response.DTO{Output: []byte(linkDontExistPhrase(link))}, nil
}

func (p *linkPresenter) UserIsNotOwnerOfLink(link string) (response.DTO, error) {
	return response.DTO{Output: []byte(userIsNotOwnerLinkPhrase(link))}, nil
}

func (p *linkPresenter) ValidationFailed(rule string) (response.DTO, error) {
	return response.DTO{Output: []byte(validationFailed(rule))}, nil
}

func NewLinkPresenter() *linkPresenter {
	return &linkPresenter{}
}
