package api

type APIResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Token  string `json:"token"`
}
