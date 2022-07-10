package dev

import (
	"context"
	"encoding/json"
	"fmt"
	"link-contractor-api/internal/controllers"
	"link-contractor-api/internal/entities/user"
	"link-contractor-api/internal/entrypoint"
	"link-contractor-api/internal/response"
	"link-contractor-api/internal/usecase/link/activatepath"
	"link-contractor-api/internal/usecase/link/create"
	"link-contractor-api/internal/usecase/link/deactivatepath"
	"link-contractor-api/internal/usecase/link/list"
	"regexp"
	"strconv"
)

type (
	ActionsPresenter interface {
		DontUnderstand() []byte
		WhatTypeLinks() []byte
	}

	Cntrls struct {
		LinkCtrl controllers.LinkController
	}

	phraseManager struct {
		actions []phraseAction
		ctrl    ctrls
	}

	phraseAction struct {
		phrasePattern  *regexp.Regexp
		makeActionFunc func(context.Context, ctrls, *regexp.Regexp, string) (entrypoint.ActionFunc, error)
	}

	ctrls struct {
		linkCtrl controllers.LinkController
	}

	payload struct {
		UserPhrase string `json:"user_phrase"`
	}
)

func (pm *phraseManager) GetAction(ctx context.Context, input []byte) (entrypoint.ActionFunc, error) {
	var pl payload
	if err := json.Unmarshal(input, &pl); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	for _, a := range pm.actions {
		if a.phrasePattern.MatchString(pl.UserPhrase) {
			return a.makeActionFunc(ctx, pm.ctrl, a.phrasePattern, pl.UserPhrase)
		}

	}

	return nil, entrypoint.UnknownActionErr
}

func New(c Cntrls, actionsPresenter ActionsPresenter) entrypoint.PhraseManager {
	pm := phraseManager{
		ctrl: ctrls{
			linkCtrl: c.LinkCtrl,
		},
		actions: initActions(actionsPresenter),
	}

	return &pm
}

func initActions(ap ActionsPresenter) []phraseAction {
	return []phraseAction{
		{
			phrasePattern: regexp.MustCompile(`^((С)|(с))оздай рандомную ссылку с длиной (?P<Length>[0-9]+), которая перебрасывает на (?P<Link>.*)$`),
			makeActionFunc: func(ctx context.Context, c ctrls, r *regexp.Regexp, inputPhrase string) (entrypoint.ActionFunc, error) {

				resMap := findGroups(inputPhrase, r, "Length", "Link")
				if resMap == nil {
					return dontUnderstand(ap), nil
				}

				link := resMap["Link"]
				lengthStr := resMap["Length"]

				length, err := strconv.ParseInt(lengthStr, 10, 32)
				if err != nil {
					return nil, fmt.Errorf("parse %s to int", lengthStr)
				}

				if link == "" || length == 0 {
					return dontUnderstand(ap), nil
				}

				return func(user user.User) (response.DTO, error) {
					linkToCreate := create.Link{
						Type:       create.Random,
						Length:     length,
						RedirectTo: link,
					}

					return c.linkCtrl.GenerateLink(ctx, linkToCreate, create.User{
						ID: user.ID,
					})
				}, nil
			},
		},
		{
			phrasePattern: regexp.MustCompile(`^((С)|(с))оздай ссылку, которая будет иметь идентификатор (?P<Path>.*) и будет вести на (?P<Link>.*)$`),
			makeActionFunc: func(ctx context.Context, c ctrls, r *regexp.Regexp, inputPhrase string) (entrypoint.ActionFunc, error) {
				resMap := findGroups(inputPhrase, r, "Path", "Link")
				if resMap == nil {
					return dontUnderstand(ap), nil
				}

				if resMap["Path"] == "" || resMap["Link"] == "" {
					return dontUnderstand(ap), nil
				}

				return func(user user.User) (response.DTO, error) {
					linkToCreate := create.Link{
						Type:          create.UserGenerated,
						RedirectTo:    resMap["Link"],
						UserGenerated: resMap["Path"],
					}

					return c.linkCtrl.GenerateLink(ctx, linkToCreate, create.User{
						ID: user.ID,
					})
				}, nil
			},
		},
		{
			phrasePattern: regexp.MustCompile(`^((Д)|(д))еактивируй ссылку с идентификатором (?P<Path>.*)$`),
			makeActionFunc: func(ctx context.Context, c ctrls, r *regexp.Regexp, inputPhrase string) (entrypoint.ActionFunc, error) {
				resMap := findGroups(inputPhrase, r, "Path")
				if resMap == nil {
					return dontUnderstand(ap), nil
				}

				if resMap["Path"] == "" {
					return dontUnderstand(ap), nil
				}

				return func(user user.User) (response.DTO, error) {
					return c.linkCtrl.DeactivatePath(ctx, deactivatepath.Path{Path: resMap["Path"]}, deactivatepath.User{ID: user.ID})
				}, nil
			},
		},

		{
			phrasePattern: regexp.MustCompile(`^((А)|(а))ктивируй ссылку с идентификатором (?P<Path>.*)$`),
			makeActionFunc: func(ctx context.Context, c ctrls, r *regexp.Regexp, inputPhrase string) (entrypoint.ActionFunc, error) {
				resMap := findGroups(inputPhrase, r, "Path")
				if resMap == nil {
					return dontUnderstand(ap), nil
				}

				if resMap["Path"] == "" {
					return dontUnderstand(ap), nil
				}

				return func(user user.User) (response.DTO, error) {
					return c.linkCtrl.ActivatePath(ctx, activatepath.Path{Path: resMap["Path"]}, activatepath.User{ID: user.ID})
				}, nil
			},
		},
		{
			phrasePattern: regexp.MustCompile(`^((П)|(п))окажи мне все мои ссылки$`),
			makeActionFunc: func(ctx context.Context, c ctrls, r *regexp.Regexp, inputPhrase string) (entrypoint.ActionFunc, error) {
				return whatTypeLinks(ap), nil
			},
		},
		{
			phrasePattern: regexp.MustCompile(`^((С)|(с)) неактивными$`),
			makeActionFunc: func(ctx context.Context, c ctrls, r *regexp.Regexp, inputPhrase string) (entrypoint.ActionFunc, error) {
				return func(user user.User) (response.DTO, error) {
					return c.linkCtrl.ListLinks(ctx, user, list.SelectOption{OnlyActive: false})
				}, nil
			},
		},
		{
			phrasePattern: regexp.MustCompile(`^((Т)|(т))олько активные$`),
			makeActionFunc: func(ctx context.Context, c ctrls, r *regexp.Regexp, inputPhrase string) (entrypoint.ActionFunc, error) {
				return func(user user.User) (response.DTO, error) {
					return c.linkCtrl.ListLinks(ctx, user, list.SelectOption{OnlyActive: true})
				}, nil
			},
		},
	}
}

func findGroups(input string, re *regexp.Regexp, groups ...string) map[string]string {
	res := re.FindAllStringSubmatch(input, -1)
	if len(res) == 0 {
		return nil
	}

	resMap := make(map[string]string, len(res))
	for _, group := range groups {
		for _, v := range res {
			for kk, vv := range re.SubexpNames() {
				if vv == group {
					resMap[group] = v[kk]
				}
			}
		}
	}

	return resMap
}

// https://tproger.ru/articles/puteshestvie-v-golang-regexp/
// хотелось бы иметь такой вот интерфейс у менеджера фраз
// Step(Phrase {
//  []On{
//		On:"bbbb",
//      Action(input) response.DTO {
//         do something
//		},
//		On:"aaaa",
//      Action(input) response.DTO {
//         do something other
//		},
//   }
//}).Step(...).Step() и т.д.

// хотя нет, надо чисто одну фразу, т.е. юзер в теории может проскочить целую ветку и придти сразу к самому последнему экшену, так будет в разы проще
