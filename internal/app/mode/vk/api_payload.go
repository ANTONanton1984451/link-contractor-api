package vk

// Event событие в вк
type Event struct {
	Type   string `json:"type"`
	Object Object `json:"object"`
}

// Object объект события, в данном случае сообщение
type Object struct {
	Message struct {
		ID       uint64 `json:"id"`
		PeerID   uint64 `json:"peer_id"`
		FromID   uint64 `json:"from_id"`
		Text     string `json:"text"`
		RandomID uint64 `json:"random_id"`
	} `json:"message"`
}

const (
	_confirmationType = "confirmation"
	_messageNew       = "message_new"
)
