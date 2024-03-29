package handlers

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	SessionId string `json:"sessionId"`
	Message   string `json:"message"`
}

type SessionResponse struct {
	Message string `json:"message"`
}
