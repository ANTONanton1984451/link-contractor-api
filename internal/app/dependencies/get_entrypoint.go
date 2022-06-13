package dependencies

import (
	devMode "link-contractor-api/internal/app/mode/dev"
	"link-contractor-api/internal/auth"
	"link-contractor-api/internal/controllers"
	"link-contractor-api/internal/dal/pool"
	"link-contractor-api/internal/dal/repository/link"
	"link-contractor-api/internal/dal/repository/user"
	"link-contractor-api/internal/entrypoint"
	devPhraseManager "link-contractor-api/internal/phrases/dev"
	"link-contractor-api/internal/presenter/dev"
	"link-contractor-api/internal/usecase/link/create"
	"link-contractor-api/internal/validation"
)

// todo переделать на геттеры для всех инстансов
func GetEntryPointForDev(config EntrypointConfig) entrypoint.Entrypoint {
	userRepo := user.New(pool.GetPool())
	linkRepo := link.New(pool.GetPool())

	authCmp := auth.New(userRepo, userRepo)
	valid := validation.New()

	createLinkUc := create.New(linkRepo, config.LinkUcRandomCreateRetries, valid)

	linkPr := dev.NewLinkPresenter()

	linkCtrl := controllers.NewLinkCtrl(linkPr, createLinkUc, config.RedirectHost)

	phManager := devPhraseManager.New(devPhraseManager.Cntrls{
		LinkCtrl: linkCtrl,
	})

	ePointPresenter := dev.NewEntryPointPresenter()
	ePoint := entrypoint.New(authCmp, phManager, ePointPresenter)

	return ePoint
}

func GetHandlerPresenter() devMode.Presenter {
	return dev.NewHandlerPresenter()
}
