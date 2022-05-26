package dev

import (
	"encoding/json"
	"fmt"
	"link-contractor-api/internal/controllers"
)

type (
	DevPresentator struct{}
)

func (p *DevPresentator) ShowGeneratedLink(link controllers.GeneratedLink) (controllers.Response, error) {
	response, err := json.Marshal(GeneratedLinkResponse{Link: link.Link})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Response: response}, nil
}

func (p *DevPresentator) ActivateOK() (controllers.Response, error) {
	response, err := json.Marshal(OkResponse{Status: _okStatus})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Response: response}, nil
}

func (p *DevPresentator) DeactivateOK() (controllers.Response, error) {
	response, err := json.Marshal(OkResponse{Status: _okStatus})
	if err != nil {
		return controllers.Response{}, fmt.Errorf("marshal: %w", err)
	}

	return controllers.Response{Response: response}, nil
}

func (p *DevPresentator) ListLinks(linkList []controllers.ListLink) (controllers.Response, error) {
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

	return controllers.Response{Response: response}, nil
}
