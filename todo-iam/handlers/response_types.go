package handlers

type RegisterResponse struct {
	LogLevel string `json:"logLevel"`
	Message  string `json:"message"`
}

type LoginResponse struct {
	Token string           `json:"token"`
	Data  RegisterResponse `json:"data"`
}
