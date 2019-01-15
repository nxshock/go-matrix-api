package matrixapi

type GuestAccess string

const (
	GuestAccessCanJoin   = "can_join"
	GuestAccessForbidden = "forbidden"
)

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id519
type MRoomGuestAccess struct {
	GuestAccess GuestAccess `json:"guest_access"` // Required. Whether guests can join the room. One of: ["can_join", "forbidden"]
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id330
type MRoomName struct {
	Name string `json:"name"` //Required. The name of the room. This MUST NOT exceed 255 bytes.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id328
type MRoomMessage struct {
	Body        string      `json:"body"`    // Required. The textual representation of this message.
	MessageType MessageType `json:"msgtype"` // Required. The type of message, e.g. m.image, m.text
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id331
type MRoomTopic struct {
	Topic string `json:"topic"` // Required. The topic text.
}

// TODO: дописать остальные элементы
