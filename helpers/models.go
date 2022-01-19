package helpers

// EncryptRequest is a struct for encrypt requests coming from the client
type EncryptRequest struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

// EncryptResponse is the response going to the client
type EncryptResponse struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

// DecryptRequest is a struct for decrypt requests coming from the client
type DecryptRequest struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

// DecryptResponse is the response going to the client
type DecryptResponse struct {
	Text string `json:"text"`
	Err  string `json:"error"`
}
