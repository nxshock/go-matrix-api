package matrixapi

import "fmt"

// APIError represents server reply in case of error
// https://matrix.org/docs/spec/client_server/r0.4.0.html#api-standards
type APIError struct {
	ErrorCode    string `json:"errcode"`
	Err          string `json:"error"`
	RetryAfterMs int    `json:"retry_after_ms"`
}

// Error return API error as string
func (apiError *APIError) Error() string {
	s := fmt.Sprintf("%s: %s", apiError.ErrorCode, apiError.Err)
	if apiError.RetryAfterMs > 0 {
		s = s + fmt.Sprintf(" (retry after %d s)", apiError.RetryAfterMs/1000)
	}

	return s
}

// ErrorCode represents Matrix error code
type ErrorCode string

const (
	MForbidden               ErrorCode = "M_FORBIDDEN"                 // Forbidden access, e.g. joining a room without permission, failed login.
	MUnknownToken            ErrorCode = "M_UNKNOWN_TOKEN"             // The access token specified was not recognised.
	MMissingToken            ErrorCode = "M_MISSING_TOKEN"             // No access token was specified for the request.
	MBadJson                 ErrorCode = "M_BAD_JSON"                  // Request contained valid JSON, but it was malformed in some way, e.g. missing required keys, invalid values for keys.
	MNotJson                 ErrorCode = "M_NOT_JSON"                  // Request did not contain valid JSON.
	MNotFound                ErrorCode = "M_NOT_FOUND"                 // No resource was found for this request.
	MLimitExceeded           ErrorCode = "M_LIMIT_EXCEEDED"            // Too many requests have been sent in a short period of time. Wait a while then try again.
	MUnknown                 ErrorCode = "M_UNKNOWN"                   // An unknown error has occurred.
	MUnrecognized            ErrorCode = "M_UNRECOGNIZED"              // The server did not understand the request.
	MUnauthorized            ErrorCode = "M_UNAUTHORIZED"              // The request was not correctly authorized. Usually due to login failures.
	MUserInUse               ErrorCode = "M_USER_IN_USE"               // Encountered when trying to register a user ID which has been taken.
	MInvalidUsername         ErrorCode = "M_INVALID_USERNAME"          // Encountered when trying to register a user ID which is not valid.
	MRoomInUse               ErrorCode = "M_ROOM_IN_USE"               // Sent when the room alias given to the createRoom API is already in use.
	MInvalidRoomState        ErrorCode = "M_INVALID_ROOM_STATE"        // Sent when the intial state given to the createRoom API is invalid.
	MThreepidInUse           ErrorCode = "M_THREEPID_IN_USE"           // Sent when a threepid given to an API cannot be used because the same threepid is already in use.
	MThreepidNotFound        ErrorCode = "M_THREEPID_NOT_FOUND"        // Sent when a threepid given to an API cannot be used because no record matching the threepid was found.
	MThreepidAuth_failed     ErrorCode = "M_THREEPID_AUTH_FAILED"      // Authentication could not be performed on the third party identifier.
	MThreepidDenied          ErrorCode = "M_THREEPID_DENIED"           // The server does not permit this third party identifier. This may happen if the server only permits, for example, email addresses from a particular domain.
	MServerNotTrusted        ErrorCode = "M_SERVER_NOT_TRUSTED"        // The client's request used a third party server, eg. identity server, that this server does not trust.
	MUnsupportedRoomVersion  ErrorCode = "M_UNSUPPORTED_ROOM_VERSION"  // The client's request to create a room used a room version that the server does not support.
	MIncompatibleRoomVersion ErrorCode = "M_INCOMPATIBLE_ROOM_VERSION" // The client attempted to join a room that has a version the server does not support. Inspect the room_version property of the error response for the room's version.
	MBadState                ErrorCode = "M_BAD_STATE"                 // The state change requested cannot be performed, such as attempting to unban a user who is not banned.
	MGuestAccessForbidden    ErrorCode = "M_GUEST_ACCESS_FORBIDDEN"    // The room or resource does not permit guests to access it.
	MCaptchaNeeded           ErrorCode = "M_CAPTCHA_NEEDED"            // A Captcha is required to complete the request.
	MCaptchaInvalid          ErrorCode = "M_CAPTCHA_INVALID"           // The Captcha provided did not match what was expected.
	MMissingParam            ErrorCode = "M_MISSING_PARAM"             // A required parameter was missing from the request.
	MInvalidParam            ErrorCode = "M_INVALID_PARAM"             // A parameter that was specified has the wrong value. For example, the server expected an integer and instead received a string.
	MTooLarge                ErrorCode = "M_TOO_LARGE"                 // The request or entity was too large.
	MExclusive               ErrorCode = "M_EXCLUSIVE"                 // The resource being requested is reserved by an application service, or the application service making the request has not created the resource.
)
