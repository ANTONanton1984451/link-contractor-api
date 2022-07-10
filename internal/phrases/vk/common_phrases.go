package vk

import (
	"link-contractor-api/internal/entities/user"
	"link-contractor-api/internal/entrypoint"
	"link-contractor-api/internal/response"
)

func dontUnderstand(aP ActionsPresenter) entrypoint.ActionFunc {
	return func(user user.User) (response.DTO, error) {
		return response.DTO{Output: aP.DontUnderstand()}, nil
	}
}

func whatTypeLinks(ap ActionsPresenter) entrypoint.ActionFunc {
	return func(user user.User) (response.DTO, error) {
		return response.DTO{Output: ap.WhatTypeLinks()}, nil
	}
}

func greeting(ap ActionsPresenter, groupLink string) entrypoint.ActionFunc {
	return func(user user.User) (response.DTO, error) {
		return response.DTO{Output: ap.Greeting(groupLink)}, nil
	}
}
