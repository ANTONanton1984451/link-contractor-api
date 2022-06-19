package controllers

import (
	"context"
	"errors"
	"link-contractor-api/internal/entities/user"
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
		ListLinks(ctx context.Context, usr user.User, option list.SelectOption) (Response, error)
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
		// todo возможно разделить на отдельные интерфейсы, пока оставим
		LinkAlreadyExist(link string) (Response, error)
		LinkIsBusy(path string) (Response, error)
		// todo возможно разделить на отдельные интерфейсы, пока оставим
		LinkDontExist(path string) (Response, error)
		UserIsNotOwnerOfLink(path string) (Response, error)
		ValidationFailed(rule string) (Response, error)
	}

	linkController struct {
		presenter LinkPresent

		createLinkUc     create.CreateLink
		activatePathUc   activatepath.ActivatePath
		deactivatePathUc deactivatepath.DeactivatePath
		listLinks        list.List

		linkDomain string
	}
)

func NewLinkCtrl(lp LinkPresent,
	createLink create.CreateLink,
	activatePath activatepath.ActivatePath,
	deactivatePath deactivatepath.DeactivatePath,
	listLinks list.List,
	linkDomain string) LinkController {

	return &linkController{
		presenter:        lp,
		createLinkUc:     createLink,
		activatePathUc:   activatePath,
		deactivatePathUc: deactivatePath,
		listLinks:        listLinks,

		linkDomain: linkDomain,
	}
}

func (ctrl *linkController) GenerateLink(ctx context.Context, link create.Link, user create.User) (Response, error) {
	path, err := ctrl.createLinkUc.Execute(ctx, link, user)
	if err != nil {
		if errors.Is(err, create.LinkAlreadyExistErr) {
			return ctrl.presenter.LinkAlreadyExist(link.RedirectTo)
		}

		if errors.Is(err, create.PathIsBusy) {
			return ctrl.presenter.LinkIsBusy(ctrl.formRedirectLink(link.UserGenerated))
		}

		var ve create.ValidateErr
		if errors.As(err, &ve) {
			return ctrl.presenter.ValidationFailed(ve.ValidateRule)
		}
		return Response{}, err
	}
	newLink := ctrl.linkDomain + "/" + path.Path

	return ctrl.presenter.ShowGeneratedLink(GeneratedLink{Link: newLink})
}

func (ctrl *linkController) ActivatePath(ctx context.Context, path activatepath.Path, user activatepath.User) (Response, error) {
	err := ctrl.activatePathUc.Execute(ctx, path, user)
	if err != nil {
		if errors.Is(err, activatepath.PathDontExistErr) {
			return ctrl.presenter.LinkDontExist(ctrl.linkDomain + path.Path)
		}

		if errors.Is(err, activatepath.UserIsNotOwnerOfPathErr) {
			return ctrl.presenter.UserIsNotOwnerOfLink(ctrl.formRedirectLink(path.Path))
		}
		return Response{}, err
	}

	return ctrl.presenter.ActivateOK()
}

func (ctrl *linkController) DeactivatePath(ctx context.Context, path deactivatepath.Path, user deactivatepath.User) (Response, error) {
	err := ctrl.deactivatePathUc.Execute(ctx, path, user)
	if err != nil {
		if errors.Is(err, deactivatepath.PathDontExistErr) {
			return ctrl.presenter.LinkDontExist(ctrl.linkDomain + path.Path)
		}

		if errors.Is(err, deactivatepath.UserIsNotOwnerOfPathErr) {
			return ctrl.presenter.UserIsNotOwnerOfLink(ctrl.formRedirectLink(path.Path))
		}
		return Response{}, err
	}

	return ctrl.presenter.DeactivateOK()
}

func (ctrl *linkController) ListLinks(ctx context.Context, usr user.User, option list.SelectOption) (Response, error) {
	links, err := ctrl.listLinks.Execute(ctx, usr, option)
	if err != nil {
		return Response{}, err
	}
	linksToPresent := make([]ListLink, 0, len(links))
	for _, l := range links {
		linksToPresent = append(linksToPresent, ListLink{
			CreatedAt: l.CreatedAt,
			From:      ctrl.formRedirectLink(l.Path),
			To:        l.RedirectTo,
			Active:    l.Active,
		})
	}
	return ctrl.presenter.ListLinks(linksToPresent)
}

func (crtl *linkController) formRedirectLink(path string) string {
	return crtl.linkDomain + "/" + path
}
