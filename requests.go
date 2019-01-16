package matrixapi

// LoginRequest represents login request
// https://matrix.org/docs/spec/client_server/r0.4.0.html#post-matrix-client-r0-login
type LoginRequest struct {
	Type                     AuthenticationType `json:"type"`                                  // Required. The login type being used. One of: ["m.login.password", "m.login.token"]
	Identifier               UserIdentifier     `json:"identifier"`                            // Identification information for the user.
	Password                 string             `json:"password,omitempty"`                    // Required when type is m.login.password. The user's password.
	Token                    string             `json:"token,omitempty"`                       // Required when type is m.login.token. Part of Token-based login.
	DeviceID                 string             `json:"device_id,omitempty"`                   // ID of the client device. If this does not correspond to a known client device, a new device will be created. The server will auto-generate a device_id if this is not specified.
	InitialDeviceDisplayName string             `json:"initial_device_display_name,omitempty"` // A display name to assign to the newly-created device. Ignored if device_id corresponds to a known device.
}

// UserIdentifier represents user identifier object
// https://matrix.org/docs/spec/client_server/r0.4.0.html#post-matrix-client-r0-login
type UserIdentifier struct {
	Type    IdentifierType `json:"type"`              // Required. The type of identification. See Identifier types for supported values and additional property descriptions.
	User    string         `json:"user,omitempty"`    // The fully qualified user ID or just local part of the user ID, to log in.
	Medium  string         `json:"medium,omitempty"`  // When logging in using a third party identifier, the medium of the identifier. Must be 'email'.
	Address string         `json:"address,omitempty"` // Third party identifier for the user.
	Country string         `json:"country,omitempty"`
	Phone   string         `json:"phone,omitempty"`
}

// CreateRoomRequest represents room creation request
// https://matrix.org/docs/spec/client_server/r0.4.0.html#post-matrix-client-r0-createroom
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
