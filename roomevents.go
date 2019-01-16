// https://matrix.org/docs/spec/client_server/r0.4.0.html#id307
package matrixapi

import (
	"encoding/json"
	"fmt"
	"time"
)

type JoinedRoomC struct {
	Name              string
	Version           string
	Creator           string
	JoinRules         JoinRule
	HistoryVisibility string
	Federate          bool
	Members           map[string]*RoomMember
	PowerLevels       MRoomPowerLevels
	GuestAccess       GuestAccess
	Topic             string
}

type InviteC struct {
	JoinRules JoinRule
	InviterID string
	Members   map[string]*RoomMember
}

func (joinedRoomC *JoinedRoomC) getOrCreateMember(memberID string) *RoomMember {
	if joinedRoomC.Members == nil {
		joinedRoomC.Members = make(map[string]*RoomMember)
	}

	if joinedRoomC.Members[memberID] == nil {
		joinedRoomC.Members[memberID] = new(RoomMember)
	}

	return joinedRoomC.Members[memberID]
}

type RoomMember struct {
	Membership  string
	DisplayName string
	AvatarURL   string
	IsDirect    bool
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id308
type MRoomAliases struct {
	// TODO: не уверен насчёт правильности
	Aliases []string `json:"aliases"` // Required. A list of room aliases.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id309
type MRoomCanonicalAlias struct {
	// TODO: не уверен насчёт правильности
	Alias string `json:"aliases"` // Required. The canonical alias.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id310
type MRoomCreate struct {
	Creator  string `json:"creator"`      // Required. The user_id of the room creator. This is set by the homeserver.
	Federate bool   `json:"m.federate"`   // Whether users on other servers can join this room. Defaults to true if key does not exist.
	Version  string `json:"room_version"` // The version of the room. Defaults to "1" if the key does not exist.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id311
type MRoomJoinRules struct {
	JoinRule JoinRule `json:"join_rule"` // Required. The type of rules used for users wishing to join this room. One of: ["public", "knock", "invite", "private"]
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id312
type MRoomMember struct {
	// TODO: уточнить структуру
	Membership  string `json:"membership"`
	DisplayName string `json:"displayname"`
	AvatarURL   string `json:"avatar_url"`
	IsDirect    bool   `json:"is_direct"`
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#m-room-power-levels
type MRoomPowerLevels struct {
	Ban           int            `json:"ban"`            // The level required to ban a user. Defaults to 50 if unspecified.
	Events        map[string]int `json:"events"`         // The level required to send specific event types. This is a mapping from event type to power level required.
	EventsDefault int            `json:"events_default"` // The default level required to send message events. Can be overridden by the events key. Defaults to 0 if unspecified.
	Invite        int            `json:"invite"`         // The level required to invite a user. Defaults to 50 if unspecified.
	Kick          int            `json:"kick"`           // The level required to kick a user. Defaults to 50 if unspecified.
	Redact        int            `json:"redact"`         // The level required to redact an event. Defaults to 50 if unspecified.
	StateDefault  int            `json:"state_default"`  // The default level required to send state events. Can be overridden by the events key. Defaults to 50 if unspecified, but 0 if there is no m.room.power_levels event at all.
	Users         map[string]int `json:"users"`          // The power levels for specific users. This is a mapping from user_id to power level for that user.
	UsersDefault  int            `json:"users_default"`  // The default power level for every user in the room, unless their user_id is mentioned in the users key. Defaults to 0 if unspecified.
	Notifications Notifications  `json:"notifications"`  // The power level requirements for specific notification types. This is a mapping from key to power level for that notifications key.
}

type Notifications struct {
	Room int `json:"room"` // The level required to trigger an @room notification. Defaults to 50 if unspecified.
}

func (client *Client) processRoomEvent(roomID string, event RoomEvent) error {
	switch event.Type {
	case "m.room.aliases":
		fmt.Println(event.Type, string(event.Content)) // TODO
	case "m.room.canonical_alias":
		fmt.Println(event.Type, string(event.Content)) // TODO
	case "m.room.create":
		var mRoomCreate MRoomCreate
		err := json.Unmarshal(event.Content, &mRoomCreate)
		if err != nil {
			return err
		}

		room := client.getOrCreateRoom(roomID)

		room.Creator = mRoomCreate.Creator
		room.Federate = mRoomCreate.Federate
		room.Version = mRoomCreate.Version
	case "m.room.join_rules":
		var mRoomJoinRules MRoomJoinRules
		err := json.Unmarshal(event.Content, &mRoomJoinRules)
		if err != nil {
			return err
		}

		client.getOrCreateRoom(roomID).JoinRules = mRoomJoinRules.JoinRule
	case "m.room.member":
		var mRoomMember MRoomMember
		err := json.Unmarshal(event.Content, &mRoomMember)
		if err != nil {
			return err
		}

		member := client.getOrCreateRoom(roomID).getOrCreateMember(event.Sender)
		member.Membership = mRoomMember.Membership
		member.DisplayName = mRoomMember.DisplayName
		member.AvatarURL = mRoomMember.AvatarURL
	case "m.room.power_levels": // TODO
		var mRoomPowerLevels MRoomPowerLevels
		err := json.Unmarshal(event.Content, &mRoomPowerLevels)
		if err != nil {
			return err
		}

		client.getOrCreateRoom(roomID).PowerLevels = mRoomPowerLevels
	case "m.room.redaction":
		fmt.Println(event.Type, string(event.Content)) // TODO
	case "m.room.guest_access":
		var mRoomGuestAccess MRoomGuestAccess
		err := json.Unmarshal(event.Content, &mRoomGuestAccess)
		if err != nil {
			return err
		}

		client.getOrCreateRoom(roomID).GuestAccess = mRoomGuestAccess.GuestAccess
	case "m.room.name":
		var mRoomName MRoomName
		err := json.Unmarshal(event.Content, &mRoomName)
		if err != nil {
			return err
		}

		client.getOrCreateRoom(roomID).Name = mRoomName.Name
	case "m.room.message":
		var mRoomMessage MRoomMessage
		err := json.Unmarshal(event.Content, &mRoomMessage)
		if err != nil {
			return err
		}

		incomingMessage := IncomingMessage{
			RoomID:   roomID,
			SenderID: event.Sender,
			Type:     mRoomMessage.MessageType,
			Body:     mRoomMessage.Body,
			Time:     time.Unix(0, event.OriginServerTs*1000000)}

		if client.MessageHandler != nil && incomingMessage.SenderID != client.UserID {
			client.MessageHandler(incomingMessage)
		}
	case "m.room.topic":
		var mRoomTopic MRoomTopic
		err := json.Unmarshal(event.Content, &mRoomTopic)
		if err != nil {
			return err
		}

		client.getOrCreateRoom(roomID).Topic = mRoomTopic.Topic
	}

	return nil
}

func (client *Client) processInviteEvent(invite *InviteC, event StrippedState) error {
	switch event.Type {
	case "m.room.join_rules":
		var mRoomJoinRules MRoomJoinRules
		err := json.Unmarshal(event.Content, &mRoomJoinRules)
		if err != nil {
			return err
		}

		invite.JoinRules = mRoomJoinRules.JoinRule
		invite.InviterID = event.Sender
	case "m.room.member":
		var mRoomMember MRoomMember
		err := json.Unmarshal(event.Content, &mRoomMember)
		if err != nil {
			return err
		}

		member := &RoomMember{
			DisplayName: mRoomMember.DisplayName,
			AvatarURL:   mRoomMember.AvatarURL,
			Membership:  mRoomMember.Membership,
			IsDirect:    mRoomMember.IsDirect} // TODO: если придёт null, сработает как false

		invite.Members[event.Sender] = member
		invite.InviterID = event.Sender
	}

	return nil
}
