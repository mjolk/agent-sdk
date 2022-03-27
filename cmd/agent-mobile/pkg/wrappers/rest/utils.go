package rest

type Response struct {
	Error   error  `json:"error"`
	Payload []byte `json:"payload"`
}
