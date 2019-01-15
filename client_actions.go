package matrixapi

import (
	"fmt"
)

func (client *Client) login(request *LoginRequest) (reply LoginReply, err error) {
	err = client.do("POST", "/_matrix/client/r0/login", request, &reply)
	return
}

/*func (client *Client) JoinedRooms() ([]string, error) {
	var joinedRoomReply JoinedRoomReply
	err := client.get("/_matrix/client/r0/joined_rooms", struct{}{}, &joinedRoomReply)
	if err != nil {
		return nil, err
	}

	return joinedRoomReply.JoinedRooms, nil
}*/

// TODO: make universal {roomId}/send/{eventType}/{txnId}
// https://matrix.org/docs/spec/client_server/r0.4.0.html#id258
func (client *Client) SendTextMessage(roomID string, text string) (eventID string, err error) {
	path := fmt.Sprintf("/_matrix/client/r0/rooms/%s/send/m.room.message", roomID)

	request := MRoomMessage{
		MessageType: MessageTypeText,
		Body:        text}

	var reply SendMessageReply

	err = client.do("POST", path, request, &reply)
	if err != nil {
		return "", err
	}

	return reply.EventID, nil
}

func (client *Client) Logout() error {
	const path = "/_matrix/client/r0/logout"

	return client.do("POST", path, struct{}{}, nil)
}

func (client *Client) CreateRoom(createRoomRequest CreateRoomRequest) (roomID string, err error) {
	const path = "/_matrix/client/r0/createRoom"

	var reply CreateRoomReply
	err = client.do("POST", path, createRoomRequest, &reply)
	if err != nil {
		return
	}

	return reply.RoomID, nil
}

func (client *Client) LeaveRoom(room string) error {
	path := fmt.Sprintf("/_matrix/client/r0/rooms/%s/leave", room)

	return client.do("POST", path, struct{}{}, nil)
}

func (client *Client) Sync() error {
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-r0-sync
	const path = "/_matrix/client/r0/sync"

	var syncRequest SyncRequest
	syncRequest.Since = client.nextBatch

	var reply SyncReply
	err := client.do("GET", path, &syncRequest, &reply)
	if err != nil {
		return err
	}
	client.nextBatch = reply.NextBatch

	for roomID, room := range reply.Rooms.Join {
		for _, event := range room.Timeline.Events {
			client.processRoomEvent(roomID, event)
		}
	}

	for roomID, invite := range reply.Rooms.Invite {
		inviteC := &InviteC{Members: make(map[string]*RoomMember)}
		client.Invites[roomID] = inviteC
		for _, event := range invite.InviteState.Events {
			client.processInviteEvent(inviteC, event)
		}
		if client.InviteHandler != nil {
			client.InviteHandler(roomID, inviteC.InviterID)
		}
	}

	return nil
}

func (client *Client) JoinRoom(roomID string) error {
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id276
	var path = fmt.Sprintf("/_matrix/client/r0/rooms/%s/join", roomID)

	var reply JoinRoomReply
	err := client.do("POST", path, struct{}{}, &reply)
	if err != nil {
		return err
	}

	// Удаляем инвайты
	delete(client.Invites, roomID)

	return nil
}
