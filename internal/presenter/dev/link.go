package dev

import (
	"encoding/json"
	"fmt"
	"link-contractor-api/internal/controllers"
)

type (
	linkPresenter struct{}
)

func (p *linkPresenter) ShowGeneratedLink(link controllers.GeneratedLink) (controllers.Response, error) {
	// todo пофиксить
	response, err := json.Marshal(OkResponse{Status: _okStatus, Result: fmt.Sprintf("Ваша ссылка - %s", link.Link)})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Output: response}, nil
}

func (p *linkPresenter) ActivateOK() (controllers.Response, error) {
	response, err := json.Marshal(OkResponse{Status: _okStatus})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Output: response}, nil
}

func (p *linkPresenter) DeactivateOK() (controllers.Response, error) {
	response, err := json.Marshal(OkResponse{Status: _okStatus})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Output: response}, nil
}

func (p *linkPresenter) ListLinks(linkList []controllers.ListLink) (controllers.Response, error) {
	links := make([]Link, 0, len(linkList))

	for _, l := range linkList {
		links = append(links, Link{
			From:      l.From,
			To:        l.To,
			Active:    l.Active,
			CreatedAt: l.CreatedAt.String(),
		})
	}
	response, err := json.Marshal(ListLinks{Links: links})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Output: response}, nil
}

func (p *linkPresenter) LinkAlreadyExist(link string) (controllers.Response, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: linkAlreadyExistPhrase(link),
	})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Output: resp}, nil
}

func (p *linkPresenter) LinkIsBusy(link string) (controllers.Response, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: linkIsBusyPhrase(link),
	})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Output: resp}, nil
}

func (p *linkPresenter) LinkDontExist(link string) (controllers.Response, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: linkDontExistPhrase(link),
	})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Output: resp}, nil
}

func (p *linkPresenter) UserIsNotOwnerOfLink(link string) (controllers.Response, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: userIsNotOwnerLink(link),
	})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Output: resp}, nil
}

func (p *linkPresenter) ValidationFailed(rule string) (controllers.Response, error) {
	resp, err := json.Marshal(FailResponse{
		Status: _failStatus,
		Reason: validationFailed(rule),
	})

	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Output: resp}, nil
}

func NewLinkPresenter() *linkPresenter {
	return &linkPresenter{}
}
