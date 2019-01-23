package matrixapi

import (
	"fmt"
)

// https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-r0-rooms-roomid-members
func (client *Client) RoomMembers(roomID string, fromCache bool) (map[string]*RoomMember, error) {
	var path = fmt.Sprintf("/_matrix/client/r0/rooms/%s/members", roomID)

	var membersReply MembersReply
	err := client.do("GET", path, nil, &membersReply)
	if err != nil {
		return nil, err
	}

	if !fromCache {
		room := client.getOrCreateRoom(roomID)
		for _, memberEvent := range membersReply.Chunk {
			member := room.Members[memberEvent.Sender]
			if member == nil {
				member = new(RoomMember)
				room.Members[memberEvent.Sender] = member // TODO: panic: assignment to entry in nil map
			}
			member.AvatarURL = memberEvent.Content.AvatarURL
			member.DisplayName = memberEvent.Content.DisplayName
			member.IsDirect = memberEvent.Content.IsDirect
			member.Membership = memberEvent.Content.Membership
		}
	}

	return client.joinedRooms[roomID].Members, nil
}
