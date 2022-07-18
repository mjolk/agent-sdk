package rest

type Response struct {
	Error   string `json:"error"`
	Payload []byte `json:"payload"`
}
