package app

import (
	"errors"
	"fmt"
	"link-contractor-api/internal/app/config"
	"link-contractor-api/internal/app/dependencies"
	"link-contractor-api/internal/app/mode/dev"
)

func Start() error {
	sa, err := getStartArgs()
	if err != nil {
		return fmt.Errorf("get work mod: %w", err)
	}

	cfg, err := config.GetConfig(sa.ConfigPath)
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}

	lg, err := createLogger(sa.Wm)
	if err != nil {
		return err
	}

	if initPoolErr := dependencies.InitPool(lg, cfg.Dev.DbDsn, cfg.Dev.MaxDbConn); initPoolErr != nil {
		return fmt.Errorf("init pool: %w", initPoolErr)
	}

	switch sa.Wm {
	case _devMod:
		if cfg.Dev == nil {
			return errors.New("empty dev config")
		}
		return dev.StartWorking(dependencies.GetEntryPointForDev(dependencies.EntrypointConfig{
			RedirectHost:              cfg.Dev.RedirectHost,
			LinkUcRandomCreateRetries: cfg.Dev.RetriesLinkCreateCount,
		}), cfg.Dev.WorkPort, lg, dependencies.GetHandlerPresenter())
	}

	return nil
}

type workMod string

var workMods = map[workMod]struct{}{
	_devMod: {},
	_vkMod:  {},
}

const (
	_devMod workMod = "dev"
	_vkMod  workMod = "vk"
)
