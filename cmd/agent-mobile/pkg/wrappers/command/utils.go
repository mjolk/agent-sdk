package command

import "bytes"

func writeResponse(w *bytes.Buffer) []byte {
	wr := &bytes.Buffer{}
	wr.WriteString("{\"payload\":")
	wr.Write(w.Bytes())
	wr.WriteString("}")
	return wr.Bytes()
}

type Response struct {
	Error   error  `json:"error"`
	Payload []byte `json:"payload"`
}
