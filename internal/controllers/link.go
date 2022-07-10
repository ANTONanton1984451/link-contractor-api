package controllers

import (
	"context"
	"errors"
	"link-contractor-api/internal/entities/user"
	"link-contractor-api/internal/response"
	"link-contractor-api/internal/usecase/link/activatepath"
	"link-contractor-api/internal/usecase/link/create"
	"link-contractor-api/internal/usecase/link/deactivatepath"
	"link-contractor-api/internal/usecase/link/list"
	"time"
)

type (
	LinkController interface {
		GenerateLink(ctx context.Context, link create.Link, user create.User) (response.DTO, error)
		ActivatePath(ctx context.Context, path activatepath.Path, user activatepath.User) (response.DTO, error)
		DeactivatePath(ctx context.Context, path deactivatepath.Path, user deactivatepath.User) (response.DTO, error)
		ListLinks(ctx context.Context, usr user.User, option list.SelectOption) (response.DTO, error)
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
		ShowGeneratedLink(link GeneratedLink) (response.DTO, error)
		ActivateOK() (response.DTO, error)
		DeactivateOK() (response.DTO, error)
		ListLinks(linkList []ListLink) (response.DTO, error)
		// todo возможно разделить на отдельные интерфейсы, пока оставим
		LinkAlreadyExist(link string) (response.DTO, error)
		LinkIsBusy(path string) (response.DTO, error)
		// todo возможно разделить на отдельные интерфейсы, пока оставим
		LinkDontExist(path string) (response.DTO, error)
		UserIsNotOwnerOfLink(path string) (response.DTO, error)
		ValidationFailed(rule string) (response.DTO, error)
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

func (ctrl *linkController) GenerateLink(ctx context.Context, link create.Link, user create.User) (response.DTO, error) {
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
		return response.DTO{}, err
	}
	newLink := ctrl.linkDomain + "/" + path.Path

	return ctrl.presenter.ShowGeneratedLink(GeneratedLink{Link: newLink})
}

func (ctrl *linkController) ActivatePath(ctx context.Context, path activatepath.Path, user activatepath.User) (response.DTO, error) {
	err := ctrl.activatePathUc.Execute(ctx, path, user)
	if err != nil {
		if errors.Is(err, activatepath.PathDontExistErr) {
			return ctrl.presenter.LinkDontExist(ctrl.linkDomain + path.Path)
		}

		if errors.Is(err, activatepath.UserIsNotOwnerOfPathErr) {
			return ctrl.presenter.UserIsNotOwnerOfLink(ctrl.formRedirectLink(path.Path))
		}
		return response.DTO{}, err
	}

	return ctrl.presenter.ActivateOK()
}

func (ctrl *linkController) DeactivatePath(ctx context.Context, path deactivatepath.Path, user deactivatepath.User) (response.DTO, error) {
	err := ctrl.deactivatePathUc.Execute(ctx, path, user)
	if err != nil {
		if errors.Is(err, deactivatepath.PathDontExistErr) {
			return ctrl.presenter.LinkDontExist(ctrl.linkDomain + path.Path)
		}

		if errors.Is(err, deactivatepath.UserIsNotOwnerOfPathErr) {
			return ctrl.presenter.UserIsNotOwnerOfLink(ctrl.formRedirectLink(path.Path))
		}
		return response.DTO{}, err
	}

	return ctrl.presenter.DeactivateOK()
}

func (ctrl *linkController) ListLinks(ctx context.Context, usr user.User, option list.SelectOption) (response.DTO, error) {
	links, err := ctrl.listLinks.Execute(ctx, usr, option)
	if err != nil {
		return response.DTO{}, err
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
