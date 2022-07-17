package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	_REDIRECT_HOST = "REDIRECT_HOST"
	_POSTGRES_USER = "POSTGRES_USER"
	_POSTGRES_DB   = "POSTGRES_DB"
	_PGSSLMODE     = "PGSSLMODE"
	_POSTGRES_HOST = "POSTGRES_HOST"
	_POSTGRES_PORT = "POSTGRES_PORT"
	_WORK_PORT     = "WORK_PORT"
	_WORK_MOD      = "WORK_MOD"
	_MAX_DB_CONN   = "MAX_DB_CONN"
	// _RETRIES_LINK_CREATE_COUNT  настройка для попытки пересоздать рандомную ссылку, если такая уже есть в базе
	_RETRIES_LINK_CREATE_COUNT        = "RETRIES_LINK_CREATE_COUNT"
	_CONNECT_DB_RETRIES               = "CONNECT_DB_RETRIES"
	_CONNECT_DB_RETRIES_WAIT_TIME_SEC = "CONNECT_DB_RETRIES_WAIT_TIME_SEC"

	// специфичные для vk конфиги
	_VK_CONFIRM_TOKEN = "VK_CONFIRM_TOKEN"
	_VK_API_URL       = "VK_API_URL"
	_VK_API_VERSION   = "VK_API_VERSION"
	_VK_ACCESS_TOKEN  = "VK_ACCESS_TOKEN"
	_VK_GROUP_LINK    = "VK_GROUP_LINK"
)

func GetConfig(configPath string) (Config, error) {
	envMap, err := godotenv.Read(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("read %s : %w", configPath, err)
	}
	cfg := Config{
		DbDsn:        formDBDsn(envMap),
		RedirectHost: envMap[_REDIRECT_HOST],
		WorkPort:     envMap[_WORK_PORT],
		WorkMod:      envMap[_WORK_MOD],

		VkConfirmToken: envMap[_VK_CONFIRM_TOKEN],
		VkAccessToken:  envMap[_VK_ACCESS_TOKEN],
		VkApiURL:       envMap[_VK_API_URL],
		VkApiVersion:   envMap[_VK_API_VERSION],
		VkGroupUrl:     envMap[_VK_GROUP_LINK],
	}

	if cfg.MaxDbConn, err = strconv.ParseInt(envMap[_MAX_DB_CONN], 10, 64); err != nil {
		return Config{}, fmt.Errorf("parse %s config to int64: %w", _MAX_DB_CONN, err)
	}

	if cfg.RetriesLinkCreateCount, err = strconv.ParseInt(envMap[_RETRIES_LINK_CREATE_COUNT], 10, 64); err != nil {
		return Config{}, fmt.Errorf("parse %s config to int64: %w", _RETRIES_LINK_CREATE_COUNT, err)
	}

	if cfg.ConnectDBRetries, err = strconv.ParseInt(envMap[_CONNECT_DB_RETRIES], 10, 64); err != nil {
		return Config{}, fmt.Errorf("parse %s config to int64: %w", _CONNECT_DB_RETRIES, err)
	}

	retryTime, err := strconv.ParseInt(envMap[_CONNECT_DB_RETRIES_WAIT_TIME_SEC], 10, 64)
	if err != nil {
		return Config{}, fmt.Errorf("parse %s config to int64: %w", _CONNECT_DB_RETRIES_WAIT_TIME_SEC, err)
	}

	cfg.ConnectDBRetriesWait = time.Duration(retryTime) * time.Second

	return cfg, nil

}

func formDBDsn(envMap map[string]string) string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s host=%s port=%s",
		envMap[_POSTGRES_USER],
		envMap[_POSTGRES_DB],
		envMap[_PGSSLMODE],
		envMap[_POSTGRES_HOST],
		envMap[_POSTGRES_PORT])
}
