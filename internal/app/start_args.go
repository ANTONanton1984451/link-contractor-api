package app

import (
	"errors"
	"flag"
	"fmt"
)

type StartArgs struct {
	Wm         workMod
	ConfigPath string
}

func getStartArgs() (StartArgs, error) {
	var (
		sa StartArgs

		mod string
	)

	flag.StringVar(&mod, "mode", string(_devMod), "Схема запуска приложения")
	flag.StringVar(&sa.ConfigPath, "config_path", "", "Путь до конфигураций")

	flag.Parse()

	if _, ok := workMods[workMod(mod)]; !ok {
		return StartArgs{}, fmt.Errorf("unknown mode %s", mod)
	}
	sa.Wm = workMod(mod)

	if sa.ConfigPath == "" {
		return StartArgs{}, errors.New("config path is empty")
	}

	return sa, nil
}
