package matrixapi

import (
	"strings"
)

func stripUsername(userFullID string) string {
	if strings.Index(userFullID, "@") != 0 {
		return ""
	}
	userFullID = userFullID[1:]

	p := strings.Index(userFullID, ":")
	if p < 0 {
		return userFullID
	}

	return userFullID[:p]
}
