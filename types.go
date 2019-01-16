package matrixapi

type MRelatesTo struct {
	InReplyTo MInReplyTo `json:"m.in_reply_to"`
}

type MInReplyTo struct {
	EventID string `json:"event_id"`
}
