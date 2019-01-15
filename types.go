package matrixapi

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id207
type IdentifierType string

const (
	IdentifierTypeUser       IdentifierType = "m.id.user"       // https://matrix.org/docs/spec/client_server/r0.4.0.html#id208
	IdentifierTypeThirdparty IdentifierType = "m.id.thirdparty" // https://matrix.org/docs/spec/client_server/r0.4.0.html#id209
	IdentifierTypePhone      IdentifierType = "m.id.phone"      // https://matrix.org/docs/spec/client_server/r0.4.0.html#id210
)

// Authentication types
// https://matrix.org/docs/spec/client_server/r0.4.0.html#id198
type AuthenticationType string

const (
	// Password-based
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id199
	AuthenticationTypePassword AuthenticationType = "m.login.password"

	// Google ReCaptcha
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id200
	AuthenticationTypeRecaptcha AuthenticationType = "m.login.recaptcha"

	// OAuth2-based
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id202
	AuthenticationTypeOauth2 AuthenticationType = "m.login.oauth2"

	// Email-based (identity server)
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id203
	AuthenticationTypeEmail AuthenticationType = "m.login.email.identity"

	// Token-based
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id201
	AuthenticationTypeToken AuthenticationType = "m.login.token"

	// Dummy Auth
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id204
	AuthenticationTypeDummy AuthenticationType = "m.login.dummy"
)

type VisibilityType string

const (
	VisibilityTypePrivate = "private"
	VisibilityTypePublic  = "public"
)
