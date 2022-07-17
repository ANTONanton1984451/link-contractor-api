package app

import (
	"errors"
	"flag"
)

// StartArgs аргументы запуска приложения
type StartArgs struct {
	ConfigPath string
}

func getStartArgs() (StartArgs, error) {
	var sa StartArgs

	flag.StringVar(&sa.ConfigPath, "config_path", "", "Путь до конфигураций")
	flag.Parse()

	if sa.ConfigPath == "" {
		return StartArgs{}, errors.New("config path is empty")
	}

	return sa, nil
}
