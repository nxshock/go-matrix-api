package matrixapi

type APIError struct {
	ErrorCode    string `json:"errcode"`
	Error        string `json:"error"`
	RetryAfterMs int    `json:"retry_after_ms"`
}
