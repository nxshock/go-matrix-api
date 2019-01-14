v - done, untested
d - deprecated

[ ]    2   API Standards
[ ]        2.1   GET /_matrix/client/versions
[ ]    3   Web Browser Clients
[ ]    4   Server Discovery
[ ]        4.1   Well-known URI
[ ]            4.1.1   GET /.well-known/matrix/client
[ ]    5   Client Authentication
[ ]        5.1   Using access tokens
[ ]        5.2   Relationship between access tokens and devices
[ ]        5.3   User-Interactive Authentication API
[ ]            5.3.1   Overview
[ ]            5.3.2   User-interactive API in the REST API
[ ]            5.3.3   Example
[ ]            5.3.4   Authentication types
[v]                5.3.4.1   Password-based
[ ]                5.3.4.2   Google ReCaptcha
[ ]                5.3.4.3   Token-based
[ ]                5.3.4.4   OAuth2-based
[ ]                5.3.4.5   Email-based (identity server)
[ ]                5.3.4.6   Dummy Auth
[ ]            5.3.5   Fallback
[ ]                5.3.5.1   Example
[ ]            5.3.6   Identifier types
[ ]                5.3.6.1   Matrix User ID
[ ]                5.3.6.2   Third-party ID
[ ]                5.3.6.3   Phone number
[ ]        5.4   Login
[v]            5.4.1   GET /_matrix/client/r0/login
[v]            5.4.2   POST /_matrix/client/r0/login
[v]            5.4.3   POST /_matrix/client/r0/logout
[ ]            5.4.4   POST /_matrix/client/r0/logout/all
[ ]            5.4.5   Login Fallback
[ ]        5.5   Account registration and management
[ ]            5.5.1   POST /_matrix/client/r0/register
[ ]            5.5.2   POST /_matrix/client/r0/register/email/requestToken
[ ]            5.5.3   POST /_matrix/client/r0/register/msisdn/requestToken
[ ]            5.5.4   POST /_matrix/client/r0/account/password
[ ]            5.5.5   POST /_matrix/client/r0/account/password/email/requestToken
[ ]            5.5.6   POST /_matrix/client/r0/account/password/msisdn/requestToken
[ ]            5.5.7   POST /_matrix/client/r0/account/deactivate
[ ]            5.5.8   GET /_matrix/client/r0/register/available
[ ]            5.5.9   Notes on password management
[ ]        5.6   Adding Account Administrative Contact Information
[ ]            5.6.1   GET /_matrix/client/r0/account/3pid
[ ]            5.6.2   POST /_matrix/client/r0/account/3pid
[ ]            5.6.3   POST /_matrix/client/r0/account/3pid/delete
[ ]            5.6.4   POST /_matrix/client/r0/account/3pid/email/requestToken
[ ]            5.6.5   POST /_matrix/client/r0/account/3pid/msisdn/requestToken
[ ]        5.7   Current account information
[ ]            5.7.1   GET /_matrix/client/r0/account/whoami
[ ]    6   Pagination
[ ]    7   Filtering
[ ]        7.1   POST /_matrix/client/r0/user/{userId}/filter
[ ]        7.2   GET /_matrix/client/r0/user/{userId}/filter/{filterId}
[ ]    8   Events
[ ]        8.1   Types of room events
[ ]        8.2   Syncing
[v]            8.2.1   GET /_matrix/client/r0/sync
[d]            8.2.2   GET /_matrix/client/r0/events
[d]            8.2.3   GET /_matrix/client/r0/initialSync
[d]            8.2.4   GET /_matrix/client/r0/events/{eventId}
[ ]        8.3   Getting events for a room
[ ]            8.3.1   GET /_matrix/client/r0/rooms/{roomId}/event/{eventId}
[ ]            8.3.2   GET /_matrix/client/r0/rooms/{roomId}/state/{eventType}/{stateKey}
[ ]            8.3.3   GET /_matrix/client/r0/rooms/{roomId}/state/{eventType}
[ ]            8.3.4   GET /_matrix/client/r0/rooms/{roomId}/state
[ ]            8.3.5   GET /_matrix/client/r0/rooms/{roomId}/members
[ ]            8.3.6   GET /_matrix/client/r0/rooms/{roomId}/joined_members
[ ]            8.3.7   GET /_matrix/client/r0/rooms/{roomId}/messages
[ ]            8.3.8   GET /_matrix/client/r0/rooms/{roomId}/initialSync
[ ]        8.4   Sending events to a room
[ ]            8.4.1   PUT /_matrix/client/r0/rooms/{roomId}/state/{eventType}/{stateKey}
[ ]            8.4.2   PUT /_matrix/client/r0/rooms/{roomId}/state/{eventType}
[ ]            8.4.3   PUT /_matrix/client/r0/rooms/{roomId}/send/{eventType}/{txnId}
[ ]        8.5   Redactions
[ ]            8.5.1   Events
[ ]                8.5.1.1   m.room.redaction
[ ]            8.5.2   Client behaviour
[ ]                8.5.2.1   PUT /_matrix/client/r0/rooms/{roomId}/redact/{eventId}/{txnId}
[ ]    9   Rooms
[ ]        9.1   Creation
[v]            9.1.1   POST /_matrix/client/r0/createRoom
[ ]        9.2   Room aliases
[ ]            9.2.1   PUT /_matrix/client/r0/directory/room/{roomAlias}
[ ]            9.2.2   GET /_matrix/client/r0/directory/room/{roomAlias}
[ ]            9.2.3   DELETE /_matrix/client/r0/directory/room/{roomAlias}
[ ]        9.3   Permissions
[ ]        9.4   Room membership
[v]            9.4.1   GET /_matrix/client/r0/joined_rooms
[ ]            9.4.2   Joining rooms
[ ]                9.4.2.1   POST /_matrix/client/r0/rooms/{roomId}/invite
[ ]                9.4.2.2   POST /_matrix/client/r0/rooms/{roomId}/join
[ ]                9.4.2.3   POST /_matrix/client/r0/join/{roomIdOrAlias}
[ ]            9.4.3   Leaving rooms
[ ]                9.4.3.1   POST /_matrix/client/r0/rooms/{roomId}/leave
[ ]                9.4.3.2   POST /_matrix/client/r0/rooms/{roomId}/forget
[ ]                9.4.3.3   POST /_matrix/client/r0/rooms/{roomId}/kick
[ ]            9.4.4   Banning users in a room
[ ]                9.4.4.1   POST /_matrix/client/r0/rooms/{roomId}/ban
[ ]                9.4.4.2   POST /_matrix/client/r0/rooms/{roomId}/unban
[ ]        9.5   Listing rooms
[ ]            9.5.1   GET /_matrix/client/r0/directory/list/room/{roomId}
[ ]            9.5.2   PUT /_matrix/client/r0/directory/list/room/{roomId}
[ ]            9.5.3   GET /_matrix/client/r0/publicRooms
[ ]            9.5.4   POST /_matrix/client/r0/publicRooms
[ ]    10   User Data
[ ]        10.1   User Directory
[ ]            10.1.1   POST /_matrix/client/r0/user_directory/search
[ ]        10.2   Profiles
[ ]            10.2.1   PUT /_matrix/client/r0/profile/{userId}/displayname
[ ]            10.2.2   GET /_matrix/client/r0/profile/{userId}/displayname
[ ]            10.2.3   PUT /_matrix/client/r0/profile/{userId}/avatar_url
[ ]            10.2.4   GET /_matrix/client/r0/profile/{userId}/avatar_url
[ ]            10.2.5   GET /_matrix/client/r0/profile/{userId}
[ ]            10.2.6   Events on Change of Profile Information
[ ]    11   Security
[ ]        11.1   Rate limiting
[ ]    12   Event Structure
[ ]        12.1   Event Fields
[ ]        12.2   Room Event Fields
[ ]        12.3   State Event Fields
[ ]        12.4   Size limits
[ ]        12.5   Room Events
[ ]            12.5.1   m.room.aliases
[ ]            12.5.2   m.room.canonical_alias
[ ]            12.5.3   m.room.create
[ ]            12.5.4   m.room.join_rules
[ ]            12.5.5   m.room.member
[ ]            12.5.6   m.room.power_levels
[ ]            12.5.7   m.room.redaction
[ ]    13   Modules
[ ]        13.1   Feature Profiles
[ ]            13.1.1   Summary
[ ]            13.1.2   Clients
[ ]                13.1.2.1   Stand-alone web (Web)
[ ]                13.1.2.2   Mobile (Mobile)
[ ]                13.1.2.3   Desktop (Desktop)
[ ]                13.1.2.4   Command Line Interface (CLI)
[ ]                13.1.2.5   Embedded (Embedded)
[ ]                    13.1.2.5.1   Application
[ ]                    13.1.2.5.2   Device
[ ]        13.2   Instant Messaging
[ ]            13.2.1   Events
[ ]                13.2.1.1   m.room.message
[ ]                13.2.1.2   m.room.message.feedback
[ ]                13.2.1.3   m.room.name
[ ]                13.2.1.4   m.room.topic
[ ]                13.2.1.5   m.room.avatar
[ ]                13.2.1.6   m.room.pinned_events
[ ]                13.2.1.7   m.room.message msgtypes
[ ]                    13.2.1.7.1   m.text
[ ]                    13.2.1.7.2   m.emote
[ ]                    13.2.1.7.3   m.notice
[ ]                    13.2.1.7.4   m.image
[ ]                    13.2.1.7.5   m.file
[ ]                    13.2.1.7.6   m.video
[ ]                    13.2.1.7.7   m.audio
[ ]                    13.2.1.7.8   m.location
[ ]            13.2.2   Client behaviour
[ ]                13.2.2.1   Recommendations when sending messages
[ ]                13.2.2.2   Local echo
[ ]                13.2.2.3   Calculating the display name for a user
[ ]                13.2.2.4   Displaying membership information with messages
[ ]                13.2.2.5   Calculating the display name for a room
[ ]                13.2.2.6   Forming relationships between events
[ ]                    13.2.2.6.1   Rich replies
[ ]                        13.2.2.6.1.1   Fallbacks and event representation
[ ]                            13.2.2.6.1.1.1   Stripping the fallback
[ ]                            13.2.2.6.1.1.2   Fallback for m.text, m.notice, and unrecognised message types
[ ]                            13.2.2.6.1.1.3   Fallback for m.emote
[ ]                            13.2.2.6.1.1.4   Fallback for m.image, m.video, m.audio, and m.file
[ ]            13.2.3   Server behaviour
[ ]            13.2.4   Security considerations
[ ]        13.3   Voice over IP
[ ]            13.3.1   Events
[ ]                13.3.1.1   m.call.invite
[ ]                13.3.1.2   m.call.candidates
[ ]                13.3.1.3   m.call.answer
[ ]                13.3.1.4   m.call.hangup
[ ]            13.3.2   Client behaviour
[ ]                13.3.2.1   Glare
[ ]            13.3.3   Server behaviour
[ ]                13.3.3.1   GET /_matrix/client/r0/voip/turnServer
[ ]            13.3.4   Security considerations
[ ]        13.4   Typing Notifications
[ ]            13.4.1   Events
[ ]                13.4.1.1   m.typing
[ ]            13.4.2   Client behaviour
[ ]                13.4.2.1   PUT /_matrix/client/r0/rooms/{roomId}/typing/{userId}
[ ]            13.4.3   Security considerations
[ ]        13.5   Receipts
[ ]            13.5.1   Events
[ ]                13.5.1.1   m.receipt
[ ]            13.5.2   Client behaviour
[ ]                13.5.2.1   POST /_matrix/client/r0/rooms/{roomId}/receipt/{receiptType}/{eventId}
[ ]            13.5.3   Server behaviour
[ ]            13.5.4   Security considerations
[ ]        13.6   Fully read markers
[ ]            13.6.1   Events
[ ]                13.6.1.1   m.fully_read
[ ]            13.6.2   Client behaviour
[ ]                13.6.2.1   POST /_matrix/client/r0/rooms/{roomId}/read_markers
[ ]            13.6.3   Server behaviour
[ ]        13.7   Presence
[ ]            13.7.1   Events
[ ]                13.7.1.1   m.presence
[ ]            13.7.2   Client behaviour
[ ]                13.7.2.1   PUT /_matrix/client/r0/presence/{userId}/status
[ ]                13.7.2.2   GET /_matrix/client/r0/presence/{userId}/status
[ ]                13.7.2.3   POST /_matrix/client/r0/presence/list/{userId}
[ ]                13.7.2.4   GET /_matrix/client/r0/presence/list/{userId}
[ ]            13.7.3   Server behaviour
[ ]                13.7.3.1   Last active ago
[ ]                13.7.3.2   Idle timeout
[ ]            13.7.4   Security considerations
[ ]        13.8   Content repository
[ ]            13.8.1   Client behaviour
[ ]                13.8.1.1   POST /_matrix/media/r0/upload
[ ]                13.8.1.2   GET /_matrix/media/r0/download/{serverName}/{mediaId}
[ ]                13.8.1.3   GET /_matrix/media/r0/download/{serverName}/{mediaId}/{fileName}
[ ]                13.8.1.4   GET /_matrix/media/r0/thumbnail/{serverName}/{mediaId}
[ ]                13.8.1.5   GET /_matrix/media/r0/preview_url
[ ]                13.8.1.6   GET /_matrix/media/r0/config
[ ]                13.8.1.7   Thumbnails
[ ]            13.8.2   Server behaviour
[ ]            13.8.3   Security considerations
[ ]        13.9   Send-to-Device messaging
[ ]            13.9.1   Client behaviour
[ ]            13.9.2   Server behaviour
[ ]            13.9.3   Protocol definitions
[ ]                13.9.3.1   PUT /_matrix/client/r0/sendToDevice/{eventType}/{txnId}
[ ]                13.9.3.2   Extensions to /sync
[ ]        13.10   Device Management
[ ]            13.10.1   Client behaviour
[ ]                13.10.1.1   GET /_matrix/client/r0/devices
[ ]                13.10.1.2   GET /_matrix/client/r0/devices/{deviceId}
[ ]                13.10.1.3   PUT /_matrix/client/r0/devices/{deviceId}
[ ]                13.10.1.4   DELETE /_matrix/client/r0/devices/{deviceId}
[ ]                13.10.1.5   POST /_matrix/client/r0/delete_devices
[ ]            13.10.2   Security considerations
[ ]        13.11   End-to-End Encryption
[ ]            13.11.1   Key Distribution
[ ]                13.11.1.1   Overview
[ ]                13.11.1.2   Key algorithms
[ ]                13.11.1.3   Device keys
[ ]                13.11.1.4   Uploading keys
[ ]                13.11.1.5   Tracking the device list for a user
[ ]                13.11.1.6   Sending encrypted attachments
[ ]                    13.11.1.6.1   Extensions to m.message msgtypes
[ ]                13.11.1.7   Claiming one-time keys
[ ]            13.11.2   Device verification
[ ]            13.11.3   Key sharing
[ ]            13.11.4   Messaging Algorithms
[ ]                13.11.4.1   Messaging Algorithm Names
[ ]                13.11.4.2   m.olm.v1.curve25519-aes-sha2
[ ]                13.11.4.3   m.megolm.v1.aes-sha2
[ ]            13.11.5   Protocol definitions
[ ]                13.11.5.1   Events
[ ]                    13.11.5.1.1   m.room.encryption
[ ]                    13.11.5.1.2   m.room.encrypted
[ ]                    13.11.5.1.3   m.room_key
[ ]                    13.11.5.1.4   m.room_key_request
[ ]                    13.11.5.1.5   m.forwarded_room_key
[ ]                13.11.5.2   Key management API
[ ]                    13.11.5.2.1   POST /_matrix/client/r0/keys/upload
[ ]                    13.11.5.2.2   POST /_matrix/client/r0/keys/query
[ ]                    13.11.5.2.3   POST /_matrix/client/r0/keys/claim
[ ]                    13.11.5.2.4   GET /_matrix/client/r0/keys/changes
[ ]                13.11.5.3   Extensions to /sync
[ ]        13.12   Room History Visibility
[ ]            13.12.1   Events
[ ]                13.12.1.1   m.room.history_visibility
[ ]            13.12.2   Client behaviour
[ ]            13.12.3   Server behaviour
[ ]            13.12.4   Security considerations
[ ]        13.13   Push Notifications
[ ]            13.13.1   Client behaviour
[ ]                13.13.1.1   GET /_matrix/client/r0/pushers
[ ]                13.13.1.2   POST /_matrix/client/r0/pushers/set
[ ]                13.13.1.3   Listing Notifications
[ ]                    13.13.1.3.1   GET /_matrix/client/r0/notifications
[ ]                13.13.1.4   Push Rules
[ ]                    13.13.1.4.1   Actions
[ ]                        13.13.1.4.1.1   Tweaks
[ ]                    13.13.1.4.2   Predefined Rules
[ ]                        13.13.1.4.2.1   Default Override Rules
[ ]                            13.13.1.4.2.1.1   .m.rule.master
[ ]                            13.13.1.4.2.1.2   .m.rule.suppress_notices
[ ]                            13.13.1.4.2.1.3   .m.rule.invite_for_me
[ ]                            13.13.1.4.2.1.4   .m.rule.member_event
[ ]                            13.13.1.4.2.1.5   .m.rule.contains_display_name
[ ]                            13.13.1.4.2.1.6   .m.rule.roomnotif
[ ]                        13.13.1.4.2.2   Default Content Rules
[ ]                            13.13.1.4.2.2.1   .m.rule.contains_user_name
[ ]                        13.13.1.4.2.3   Default Underride Rules
[ ]                            13.13.1.4.2.3.1   .m.rule.call
[ ]                            13.13.1.4.2.3.2   .m.rule.encrypted_room_one_to_one
[ ]                            13.13.1.4.2.3.3   .m.rule.room_one_to_one
[ ]                            13.13.1.4.2.3.4   .m.rule.message
[ ]                            13.13.1.4.2.3.5   .m.rule.encrypted
[ ]                    13.13.1.4.3   Conditions
[ ]                13.13.1.5   Push Rules: API
[ ]                    13.13.1.5.1   GET /_matrix/client/r0/pushrules/
[ ]                    13.13.1.5.2   GET /_matrix/client/r0/pushrules/{scope}/{kind}/{ruleId}
[ ]                    13.13.1.5.3   DELETE /_matrix/client/r0/pushrules/{scope}/{kind}/{ruleId}
[ ]                    13.13.1.5.4   PUT /_matrix/client/r0/pushrules/{scope}/{kind}/{ruleId}
[ ]                    13.13.1.5.5   GET /_matrix/client/r0/pushrules/{scope}/{kind}/{ruleId}/enabled
[ ]                    13.13.1.5.6   PUT /_matrix/client/r0/pushrules/{scope}/{kind}/{ruleId}/enabled
[ ]                    13.13.1.5.7   GET /_matrix/client/r0/pushrules/{scope}/{kind}/{ruleId}/actions
[ ]                    13.13.1.5.8   PUT /_matrix/client/r0/pushrules/{scope}/{kind}/{ruleId}/actions
[ ]                13.13.1.6   Push Rules: Events
[ ]                    13.13.1.6.1   Examples
[ ]            13.13.2   Server behaviour
[ ]            13.13.3   Push Gateway behaviour
[ ]                13.13.3.1   Recommendations for APNS
[ ]            13.13.4   Security considerations
[ ]        13.14   Third party invites
[ ]            13.14.1   Events
[ ]                13.14.1.1   m.room.third_party_invite
[ ]            13.14.2   Client behaviour
[ ]                13.14.2.1   POST /_matrix/client/r0/rooms/{roomId}/invite
[ ]            13.14.3   Server behaviour
[ ]            13.14.4   Security considerations
[ ]        13.15   Server Side Search
[ ]            13.15.1   Client behaviour
[ ]                13.15.1.1   POST /_matrix/client/r0/search
[ ]            13.15.2   Search Categories
[ ]                13.15.2.1   room_events
[ ]            13.15.3   Ordering
[ ]            13.15.4   Groups
[ ]            13.15.5   Pagination
[ ]            13.15.6   Security considerations
[ ]        13.16   Guest Access
[ ]            13.16.1   Events
[ ]                13.16.1.1   m.room.guest_access
[ ]            13.16.2   Client behaviour
[ ]            13.16.3   Server behaviour
[ ]            13.16.4   Security considerations
[ ]        13.17   Room Previews
[ ]            13.17.1   Client behaviour
[ ]                13.17.1.1   GET /_matrix/client/r0/events
[ ]            13.17.2   Server behaviour
[ ]            13.17.3   Security considerations
[ ]        13.18   Room Tagging
[ ]            13.18.1   Events
[ ]                13.18.1.1   m.tag
[ ]            13.18.2   Client Behaviour
[ ]                13.18.2.1   GET /_matrix/client/r0/user/{userId}/rooms/{roomId}/tags
[ ]                13.18.2.2   PUT /_matrix/client/r0/user/{userId}/rooms/{roomId}/tags/{tag}
[ ]                13.18.2.3   DELETE /_matrix/client/r0/user/{userId}/rooms/{roomId}/tags/{tag}
[ ]        13.19   Client Config
[ ]            13.19.1   Events
[ ]            13.19.2   Client Behaviour
[ ]                13.19.2.1   PUT /_matrix/client/r0/user/{userId}/account_data/{type}
[ ]                13.19.2.2   PUT /_matrix/client/r0/user/{userId}/rooms/{roomId}/account_data/{type}
[ ]            13.19.3   Server Behaviour
[ ]        13.20   Server Administration
[ ]            13.20.1   Client Behaviour
[ ]                13.20.1.1   GET /_matrix/client/r0/admin/whois/{userId}
[ ]        13.21   Event Context
[ ]            13.21.1   Client behaviour
[ ]                13.21.1.1   GET /_matrix/client/r0/rooms/{roomId}/context/{eventId}
[ ]            13.21.2   Security considerations
[ ]        13.22   CAS-based client login
[ ]            13.22.1   Client behaviour
[ ]                13.22.1.1   GET /_matrix/client/r0/login/cas/redirect
[ ]                13.22.1.2   GET /_matrix/client/r0/login/cas/ticket
[ ]            13.22.2   Server behaviour
[ ]                13.22.2.1   Handling the redirect endpoint
[ ]                13.22.2.2   Handling the authentication endpoint
[ ]        13.23   Direct Messaging
[ ]            13.23.1   Events
[ ]                13.23.1.1   m.direct
[ ]            13.23.2   Client behaviour
[ ]            13.23.3   Server behaviour
[ ]        13.24   Ignoring Users
[ ]            13.24.1   Events
[ ]                13.24.1.1   m.ignored_user_list
[ ]            13.24.2   Client behaviour
[ ]            13.24.3   Server behaviour
[ ]        13.25   Sticker Messages
[ ]            13.25.1   Events
[ ]                13.25.1.1   m.sticker
[ ]            13.25.2   Client behaviour
[ ]        13.26   Reporting Content
[ ]            13.26.1   Client behaviour
[ ]                13.26.1.1   POST /_matrix/client/r0/rooms/{roomId}/report/{eventId}
[ ]            13.26.2   Server behaviour
[ ]        13.27   Third Party Networks
[ ]            13.27.1   Third Party Lookups
[ ]                13.27.1.1   GET /_matrix/client/r0/thirdparty/protocols
[ ]                13.27.1.2   GET /_matrix/client/r0/thirdparty/protocol/{protocol}
[ ]                13.27.1.3   GET /_matrix/client/r0/thirdparty/location/{protocol}
[ ]                13.27.1.4   GET /_matrix/client/r0/thirdparty/user/{protocol}
[ ]                13.27.1.5   GET /_matrix/client/r0/thirdparty/location
[ ]                13.27.1.6   GET /_matrix/client/r0/thirdparty/user
[ ]        13.28   OpenID
[ ]            13.28.1   POST /_matrix/client/r0/user/{userId}/openid/request_token
[ ]        13.29   Server Access Control Lists (ACLs) for rooms
[ ]            13.29.1   m.room.server_acl
[ ]            13.29.2   Client behaviour
[ ]            13.29.3   Server behaviour
[ ]            13.29.4   Security considerations
[ ]        13.30   User, room, and group mentions
[ ]            13.30.1   Client behaviour
