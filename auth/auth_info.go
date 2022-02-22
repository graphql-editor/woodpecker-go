package auth

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// APIKey implements authentication using API Key on Basic Auth
type APIKey string

// AuthenticateRequest implements  runtime.ClientAuthInfoWriter
func (a APIKey) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	req.SetHeaderParam("Authentication", "Basic "+string(a))
	return nil
}
