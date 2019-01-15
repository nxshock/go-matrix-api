package matrixapi

// Invite3pid represents third party IDs to invite into the room
// https://matrix.org/docs/spec/client_server/r0.4.0.html#post-matrix-client-r0-createroom
type Invite3pid struct {
	IDServer string `json:"id_server"` // Required. The hostname+port of the identity server which should be used for third party identifier lookups.
	Medium   string `json:"medium"`    // Required. The kind of address being passed in the address field, for example email.
	Address  string `json:"address"`   // Required. The invitee's third party identifier.
}
