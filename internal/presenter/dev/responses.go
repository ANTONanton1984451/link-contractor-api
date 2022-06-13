package dev

type (
	OkResponse struct {
		Status string `json:"status"`
		Result string `json:"result"`
	}

	FailResponse struct {
		Status string `json:"status"`
		Reason string `json:"reason"`
	}

	ListLinks struct {
		Links []Link `json:"links"`
	}

	Link struct {
		From      string `json:"from"`
		To        string `json:"to"`
		Active    bool   `json:"active"`
		CreatedAt string `json:"createdAt"`
	}
)

const (
	_okStatus   = "OK"
	_failStatus = "FAIL"
)
