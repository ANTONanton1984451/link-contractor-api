package dev

type (
	GeneratedLinkResponse struct {
		Link string `json:"link"`
	}

	OkResponse struct {
		Status string `json:"status"`
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
	_okStatus = "OK"
)
