package response

type Response struct {
	Message string   `json:"message,omitempty"`
	Data    any      `json:"data,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}
