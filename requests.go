package matrixapi

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id213
type LoginRequest struct {
	Type                     AuthenticationType `json:"type"`
	User                     string             `json:"user"`
	Password                 string             `json:"password"`
	Token                    string             `json:"token,omitempty"`
	DeviceID                 string             `json:"device_id,omitempty"`
	InitialDeviceDisplayName string             `json:"initial_device_display_name,omitempty"`
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id264
type CreateRoomRequest struct {
	Visibility    VisibilityType `json:"visibility,omitempty"`
	RoomAliasName string         `json:"room_alias_name,omitempty"`
	Name          string         `json:"name,omitempty"`
	Topic         string         `json:"topic,omitempty"`
	Invite        []string       `json:"invite,omitempty"`
	Invite3pids   []Invite3pid   `json:"invite_3pid,omitempty"`
	RoomVersion   string         `json:"room_version,omitempty"`
	// TODO: проверить тип
	// CreationContent CreationContentType `json:"creation_content,omitempty"`
	InitialState []StateEvent `json:"initial_state,omitempty"`
	Preset       string       `json:"preset,omitempty"` // TODO: проверить тип
	IsDirect     bool         `json:"is_direct,omitempty"`
	// PowerLevelContentOverride `json:"power_level_content_override"`
}
