package structures

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Jwt     string      `json:"jwt"`
	Data    interface{} `json:"data"`
}
type ErrorResponse struct {
	Error   string `json:"error"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

type AuthErrorResponse struct {
	Message string `json:"message"`
}
