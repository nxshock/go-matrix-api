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
