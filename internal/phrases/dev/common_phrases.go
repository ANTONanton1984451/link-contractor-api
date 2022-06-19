package dev

import (
	"link-contractor-api/internal/controllers"
	"link-contractor-api/internal/entities/user"
	"link-contractor-api/internal/entrypoint"
)

func dontUnderstand(aP ActionsPresenter) entrypoint.ActionFunc {
	return func(user user.User) (controllers.Response, error) {
		return controllers.Response{Output: aP.DontUnderstand()}, nil
	}
}

func whatTypeLinks(ap ActionsPresenter) entrypoint.ActionFunc {
	return func(user user.User) (controllers.Response, error) {
		return controllers.Response{Output: ap.WhatTypeLinks()}, nil
	}
}
