package app

import (
	"fmt"
	"link-contractor-api/internal/app/dependencies"
	"link-contractor-api/internal/app/mode/dev"
	"link-contractor-api/internal/app/mode/vk"
	"link-contractor-api/internal/config"
)

// Start старт приложения
func Start() error {
	sa, err := getStartArgs()
	if err != nil {
		return fmt.Errorf("get work mod: %w", err)
	}

	cfg, err := config.GetConfig(sa.ConfigPath)
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}

	lg, err := createLogger(workMod(cfg.WorkMod))
	if err != nil {
		return err
	}

	if initPoolErr := dependencies.InitPool(lg, cfg.DbDsn, cfg.MaxDbConn, cfg.ConnectDBRetries, cfg.ConnectDBRetriesWait); initPoolErr != nil {
		return fmt.Errorf("init pool: %w", initPoolErr)
	}

	switch workMod(cfg.WorkMod) {
	case _devMod:
		return dev.StartWorking(dependencies.GetEntrypointForDev(dependencies.EntrypointConfig{
			RedirectHost:              cfg.RedirectHost,
			LinkUcRandomCreateRetries: cfg.RetriesLinkCreateCount,
		}), cfg.WorkPort, lg, dependencies.GetHandlerPresenterForDev())
	case _vkMod:
		workModConfig := vk.Config{
			WorkPort:     cfg.WorkPort,
			ApiURL:       cfg.VkApiURL,
			ApiVersion:   cfg.VkApiVersion,
			AccessToken:  cfg.VkAccessToken,
			ConfirmToken: cfg.VkConfirmToken,
		}

		epConfig := dependencies.EntrypointConfig{
			RedirectHost:              cfg.RedirectHost,
			LinkUcRandomCreateRetries: cfg.RetriesLinkCreateCount,

			VkGroupUrl: cfg.VkGroupUrl,
		}

		return vk.StartWorking(dependencies.GetEntrypointForVk(epConfig),
			workModConfig,
			lg,
			dependencies.GetHandlerPresenterForVk(),
			dependencies.GetLinkRepository())
	default:
		return fmt.Errorf("unknown work mod %s", cfg.WorkMod)
	}

}

type workMod string

const (
	_devMod workMod = "dev"
	_vkMod  workMod = "vk"
)
