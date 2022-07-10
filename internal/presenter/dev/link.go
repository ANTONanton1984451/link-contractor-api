package dev

import (
	"encoding/json"
	"fmt"
	"link-contractor-api/internal/controllers"
	"link-contractor-api/internal/response"
)

type (
	linkPresenter struct{}
)

func (p *linkPresenter) ShowGeneratedLink(link controllers.GeneratedLink) (response.DTO, error) {

	resp, err := json.Marshal(OkResponse{Status: _okStatus, Result: fmt.Sprintf("Ваша ссылка - %s", link.Link)})
	if err != nil {
		return response.DTO{}, fmt.Errorf("marshal: %w", err)
	}

	return response.DTO{Output: resp}, nil
}

func (p *linkPresenter) ActivateOK() (response.DTO, error) {
	resp, err := json.Marshal(OkResponse{Status: _okStatus, Result: `ссылка активирована`})
	if err != nil {
		return response.DTO{}, fmt.Errorf("marshal: %w", err)
	}

	return response.DTO{Output: resp}, nil
}

func (p *linkPresenter) DeactivateOK() (response.DTO, error) {
	resp, err := json.Marshal(OkResponse{Status: _okStatus, Result: `ссылка деактивирована`})
	if err != nil {
		return response.DTO{}, fmt.Errorf("marshal: %w", err)
	}

	return response.DTO{Output: resp}, nil
}

func (p *linkPresenter) ListLinks(linkList []controllers.ListLink) (response.DTO, error) {
	if len(linkList) == 0 {
		resp, _ := json.Marshal(OkResponse{Status: _okStatus, Result: `у тебя нет ссылок на данный момент`})
		return response.DTO{Output: resp}, nil
	}
	links := make([]Link, 0, len(linkList))

	for _, l := range linkList {
		links = append(links, Link{
			From:      l.From,
			To:        l.To,
			Active:    l.Active,
			CreatedAt: l.CreatedAt.String(),
		})
	}
	resp, err := json.Marshal(ListLinks{Links: links})
	if err != nil {
		return response.DTO{}, fmt.Errorf("marshal: %w", err)
	}

	return response.DTO{Output: resp}, nil
}

func (p *linkPresenter) LinkAlreadyExist(link string) (response.DTO, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: linkAlreadyExistPhrase(link),
	})
	if err != nil {
		return response.DTO{}, fmt.Errorf("marshal: %w", err)
	}

	return response.DTO{Output: resp}, nil
}

func (p *linkPresenter) LinkIsBusy(link string) (response.DTO, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: linkIsBusyPhrase(link),
	})
	if err != nil {
		return response.DTO{}, fmt.Errorf("marshal: %w", err)
	}

	return response.DTO{Output: resp}, nil
}

func (p *linkPresenter) LinkDontExist(link string) (response.DTO, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: linkDontExistPhrase(link),
	})
	if err != nil {
		return response.DTO{}, fmt.Errorf("marshal: %w", err)
	}

	return response.DTO{Output: resp}, nil
}

func (p *linkPresenter) UserIsNotOwnerOfLink(link string) (response.DTO, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: userIsNotOwnerLink(link),
	})
	if err != nil {
		return response.DTO{}, fmt.Errorf("marshal: %w", err)
	}

	return response.DTO{Output: resp}, nil
}

func (p *linkPresenter) ValidationFailed(rule string) (response.DTO, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: validationFailed(rule),
	})

	if err != nil {
		return response.DTO{}, fmt.Errorf("marshal: %w", err)
	}

	return response.DTO{Output: resp}, nil
}

func NewLinkPresenter() *linkPresenter {
	return &linkPresenter{}
}
