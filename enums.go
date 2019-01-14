package matrixapi

type Membership string

const (
	MembershipInvite Membership = "invite"
	MembershipJoin   Membership = "join"
	MembershipKnock  Membership = "knock"
	MembershipLeave  Membership = "leave"
	MembershipBan    Membership = "ban"
)

type SetPresence string

const (
	SetPresenceOffline     SetPresence = "offline"
	SetPresenceOnline      SetPresence = "online"
	SetPresenceUnavailable SetPresence = "unavailable"
)

type JoinRule string

const (
	JoinRulePublic  JoinRule = "public"
	JoinRuleKnock   JoinRule = "knock"
	JoinRuleInvite  JoinRule = "invite"
	JoinRulePrivate JoinRule = "private"
)
