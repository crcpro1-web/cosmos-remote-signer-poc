package main

type Request struct {
	Type         string `json:"type"`
	SignBytes    string `json:"sign_bytes,omitempty"`
	VoteExtBytes string `json:"vote_ext_bytes,omitempty"`

	Height int64  `json:"height,omitempty"`
	Round  int32  `json:"round,omitempty"`
	Step   string `json:"step,omitempty"`
}

type Response struct {
	Status             string `json:"status"`
	Signature          string `json:"signature,omitempty"`
	ExtensionSignature string `json:"extension_signature,omitempty"`
	Error              string `json:"error,omitempty"`
}
