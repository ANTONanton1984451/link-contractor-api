package controllers

import (
	"context"
	"link-contractor-api/internal/usecase/link/activatepath"
	"link-contractor-api/internal/usecase/link/create"
	"link-contractor-api/internal/usecase/link/deactivatepath"
	"link-contractor-api/internal/usecase/link/list"
	"time"
)

type (
	LinkController interface {
		GenerateLink(ctx context.Context, link create.Link, user create.User) (Response, error)
		ActivatePath(ctx context.Context, path activatepath.Path, user activatepath.User) (Response, error)
		DeactivatePath(ctx context.Context, path deactivatepath.Path, user deactivatepath.User) (Response, error)
		ListLinks(ctx context.Context, option list.SelectOption) (Response, error)
	}
	GeneratedLink struct {
		Link string
	}

	ListLink struct {
		From      string
		To        string
		CreatedAt time.Time
		Active    bool
	}

	LinkPresent interface {
		ShowGeneratedLink(link GeneratedLink) (Response, error)
		ActivateOK() (Response, error)
		DeactivateOK() (Response, error)
		ListLinks(linkList []ListLink) (Response, error)
	}

	linkController struct {
		presenter LinkPresent

		createLinkUc   create.CreateLink
		activatepath   activatepath.ActivatePath
		deactivatepath deactivatepath.DeactivatePath
		listLinks      list.List

		linkDomain string
	}
)

func (ctrl *linkController) GenerateLink(ctx context.Context, link create.Link, user create.User) (Response, error) {
	path, err := ctrl.createLinkUc.Execute(ctx, link, user)
	if err != nil {
		return Response{}, err
	}
	newLink := ctrl.linkDomain + "/" + path.Path

	return ctrl.presenter.ShowGeneratedLink(GeneratedLink{Link: newLink})
}

func (ctrl *linkController) ActivatePath(ctx context.Context, path activatepath.Path, user activatepath.User) (Response, error) {
	err := ctrl.activatepath.Execute(ctx, path, user)
	if err != nil {
		return Response{}, err
	}

	return ctrl.presenter.ActivateOK()
}

func (ctrl *linkController) DeactivatePath(ctx context.Context, path deactivatepath.Path, user deactivatepath.User) (Response, error) {
	err := ctrl.deactivatepath.Execute(ctx, path, user)
	if err != nil {
		return Response{}, err
	}

	return ctrl.presenter.DeactivateOK()
}

func (ctrl *linkController) ListLinks(ctx context.Context, option list.SelectOption) (Response, error) {
	links, err := ctrl.listLinks.Execute(ctx, option)
	if err != nil {
		return Response{}, err
	}
	linksToPresent := make([]ListLink, 0, len(links))
	for _, l := range links {
		linksToPresent = append(linksToPresent, ListLink{
			CreatedAt: l.CreatedAt,
			From:      ctrl.linkDomain + "/" + l.Path,
			To:        l.RedirectTo,
			Active:    l.Active,
		})
	}
	return ctrl.presenter.ListLinks(linksToPresent)
}
