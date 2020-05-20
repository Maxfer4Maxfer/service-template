package http

// Response represent a struct to pass to a http response.
type Response struct {
	Code    int         `json:"code"`
	Content interface{} `json:"content"`
	Error   string      `json:"error,omitempty"`
}
