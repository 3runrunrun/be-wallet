package account

// BasicResponse for http response
type BasicResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
