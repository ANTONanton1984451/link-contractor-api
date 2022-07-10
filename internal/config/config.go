package config

import "time"

type (
	Config struct {
		DbDsn                  string
		RedirectHost           string
		WorkPort               string
		MaxDbConn              int64
		RetriesLinkCreateCount int64
		ConnectDBRetries       int64
		ConnectDBRetriesWait   time.Duration
		WorkMod                string

		VkConfirmToken string
		VkApiURL       string
		VkApiVersion   string
		VkAccessToken  string
		VkGroupUrl     string
	}
)
