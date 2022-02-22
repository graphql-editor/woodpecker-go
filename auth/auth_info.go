package auth

import (
	"encoding/base64"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// APIKey implements authentication using API Key on Basic Auth. Expected Base64
type APIKey string

// NewAPIKey encodes api key with base64 and returns APIKey for woodpecker
func NewAPIKey(apiKey string) APIKey {
	return APIKey(base64.StdEncoding.EncodeToString([]byte(apiKey)))
}

// AuthenticateRequest implements  runtime.ClientAuthInfoWriter
func (a APIKey) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	req.SetHeaderParam("Authorization", "Basic "+string(a))
	return nil
}
