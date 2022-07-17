package dependencies

import (
	devMod "link-contractor-api/internal/app/mode/dev"
	vkMod "link-contractor-api/internal/app/mode/vk"
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
	"link-contractor-api/internal/usecase/link/activatepath"
	"link-contractor-api/internal/usecase/link/create"
	"link-contractor-api/internal/usecase/link/deactivatepath"
	"link-contractor-api/internal/usecase/link/list"
	"link-contractor-api/internal/validation"
)

// todo здесь можно сделать DI контейнер, но т.к. компонентов мало, то я не буду так делать

// GetEntrypointForDev получение Entrypoint для разработки
func GetEntrypointForDev(config EntrypointConfig) entrypoint.Entrypoint {
	userRepo := user.New(pool.GetPool())
	linkRepo := link.New(pool.GetPool())

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

// GetEntrypointForVk получение Entrypoint для работы в вк
func GetEntrypointForVk(config EntrypointConfig) entrypoint.Entrypoint {
	userRepo := user.New(pool.GetPool())
	linkRepo := link.New(pool.GetPool())

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

// GetHandlerPresenterForDev презентор для дев режима работы
func GetHandlerPresenterForDev() devMod.Presenter {
	return dev.NewHandlerPresenter()
}

// GetHandlerPresenterForVk презентор для режима работы для вк
func GetHandlerPresenterForVk() vkMod.Presenter {
	return vk.NewHandlerPresenter()
}
