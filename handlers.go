package matrixapi

import (
	"time"
)

type InviteHandler func(roomID, inviterID string)
type MessageHandler func(IncomingMessage)

type IncomingMessage struct {
	RoomID   string
	SenderID string
	Type     MessageType
	Body     string
	Time     time.Time
}
