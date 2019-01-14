package matrixapi

type SendMessageRequest struct {
	MessageType MessageType `json:"msgtype,omitempty"`
	Body        string      `json:"body"`
}
