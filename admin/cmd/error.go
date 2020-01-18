package cmd

type ErrorInfo struct {
	Error string `json:"error"`
}

func DecodeError(data []byte) (string, error) {

	return string(data), nil
}
func EncodeError(err string) ([]byte, error) {

	return []byte(err), nil
}
