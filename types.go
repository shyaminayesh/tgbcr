package tgbcr

type From struct {
	ID uint64 `json:"id"`
}

type Chat struct {
	ID uint64 `json:"id"`
}

type Message struct {
	ID   uint64 `json:"message_id"`
	Date uint64 `json:"date"`
	Text string `json:"text"`

	From From `json:"from"`
	Chat Chat `json:"chat"`
}

type Context struct {
	Message Message `json:"message"`
}
