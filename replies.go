package matrixapi

type JoinedRoomReply struct {
	JoinedRooms []string `json:"joined_rooms"`
}

// SendMessageReply represents reply for send message command
// https://matrix.org/docs/spec/client_server/r0.4.0.html#put-matrix-client-r0-rooms-roomid-state-eventtype-statekey
type SendMessageReply struct {
	EventID string `json:"event_id"` // A unique identifier for the event.
}

type CreateRoomReply struct {
	RoomID string `json:"room_id"`
}

type GetLoginReply struct {
	Flows []Flow // The homeserver's supported login types
}

type LoginReply struct {
	AccessToken string `json:"access_token"`
	HomeServer  string `json:"home_server"`
	UserID      string `json:"user_id"`
}

type SyncReply struct {
	NextBatch              string         `json:"next_batch"`                 // Required. The batch token to supply in the since param of the next /sync request.
	Rooms                  RoomsSyncReply `json:"rooms"`                      // Updates to rooms.
	Presence               Presence       `json:"presence"`                   // The updates to the presence status of other users.
	AccountData            AccountData    `json:"account_data"`               // The global private data created by this user.
	ToDevice               ToDevice       `json:"to_device"`                  // Information on the send-to-device messages for the client device, as defined in Send-to-Device messaging.
	DeviceLists            DeviceLists    `json:"device_lists"`               // Information on end-to-end device updates, as specified in End-to-end encryption.
	DeviceOneTimeKeysCount map[string]int `json:"device_one_time_keys_count"` // Information on end-to-end encryption keys, as specified in End-to-end encryption.
}

type RoomsSyncReply struct {
	Join   map[string]JoinedRoom  `json:"join"`   // The rooms that the user has joined.
	Invite map[string]InvitedRoom `json:"invite"` // The rooms that the user has been invited to.
	Leave  map[string]LeftRoom    `json:"leave"`  // The rooms that the user has left or been banned from.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id276
type JoinRoomReply struct {
	RoomID string `json:"room_id"` // The joined room ID must be returned in the room_id field.
}
