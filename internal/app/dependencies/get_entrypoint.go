package dependencies

import (
	devMod "link-contractor-api/internal/app/mode/dev"
	"link-contractor-api/internal/auth"
	"link-contractor-api/internal/controllers"
	"link-contractor-api/internal/dal/pool"
	"link-contractor-api/internal/dal/repository/link"
	"link-contractor-api/internal/dal/repository/user"
	"link-contractor-api/internal/entrypoint"
	devPhraseManager "link-contractor-api/internal/phrases/dev"
	vkPhraseManager "link-contractor-api/internal/phrases/vk"
	"link-contractor-api/internal/presenter/dev"
	"link-contractor-api/internal/presenter/vk"
	"link-contractor-api/internal/redirect"
	"link-contractor-api/internal/usecase/link/activatepath"
	"link-contractor-api/internal/usecase/link/create"
	"link-contractor-api/internal/usecase/link/deactivatepath"
	"link-contractor-api/internal/usecase/link/list"
	"link-contractor-api/internal/validation"
)

// todo здесь можно сделать DI контейнер, но т.к. компонентов мало, то я не буду так делать
func GetEntryPointForDev(config EntrypointConfig) entrypoint.Entrypoint {
	rd := redirect.New()

	userRepo := user.New(pool.GetPool())
	linkRepo := link.New(pool.GetPool(), rd)

	linkPr := dev.NewLinkPresenter()
	ePointPresenter := dev.NewEntryPointPresenter()

	authCmp := auth.New(userRepo, userRepo)
	valid := validation.New()

	createLinkUc := create.New(linkRepo, config.LinkUcRandomCreateRetries, valid)
	activatePath := activatepath.New(linkRepo)
	deactivatePath := deactivatepath.New(linkRepo)
	listLinks := list.New(linkRepo)

	phManagerPresenter := dev.NewPhraseManagerPresenter()

	linkCtrl := controllers.NewLinkCtrl(linkPr, createLinkUc, activatePath, deactivatePath, listLinks, config.RedirectHost)

	phManager := devPhraseManager.New(devPhraseManager.Cntrls{LinkCtrl: linkCtrl}, phManagerPresenter)

	ePoint := entrypoint.New(authCmp, phManager, ePointPresenter)

	return ePoint
}

func GetEntryPointForVk(config EntrypointConfig) entrypoint.Entrypoint {
	rd := redirect.New()

	userRepo := user.New(pool.GetPool())
	linkRepo := link.New(pool.GetPool(), rd)

	linkPr := vk.NewLinkPresenter()
	ePointPresenter := vk.NewEntryPointPresenter()

	authCmp := auth.New(userRepo, userRepo)
	valid := validation.New()

	createLinkUc := create.New(linkRepo, config.LinkUcRandomCreateRetries, valid)
	activatePath := activatepath.New(linkRepo)
	deactivatePath := deactivatepath.New(linkRepo)
	listLinks := list.New(linkRepo)

	phManagerPresenter := vk.NewPhraseManagerPresenter()

	linkCtrl := controllers.NewLinkCtrl(linkPr, createLinkUc, activatePath, deactivatePath, listLinks, config.RedirectHost)

	phManager := vkPhraseManager.New(vkPhraseManager.Cntrls{LinkCtrl: linkCtrl}, phManagerPresenter, config.VkGroupUrl)

	ePoint := entrypoint.New(authCmp, phManager, ePointPresenter)

	return ePoint
}

func GetHandlerPresenter() devMod.Presenter {
	return dev.NewHandlerPresenter()
}
