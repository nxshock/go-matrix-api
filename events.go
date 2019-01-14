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

// TODO: дописать остальные элементы
