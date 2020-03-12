// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OptionalClaim undocumented
type OptionalClaim struct {
	// Object is the base model of OptionalClaim
	Object
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Source undocumented
	Source *string `json:"source,omitempty"`
	// Essential undocumented
	Essential *bool `json:"essential,omitempty"`
	// AdditionalProperties undocumented
	AdditionalProperties []string `json:"additionalProperties,omitempty"`
}

// OptionalClaims undocumented
type OptionalClaims struct {
	// Object is the base model of OptionalClaims
	Object
	// IDToken undocumented
	IDToken []OptionalClaim `json:"idToken,omitempty"`
	// AccessToken undocumented
	AccessToken []OptionalClaim `json:"accessToken,omitempty"`
	// Saml2Token undocumented
	Saml2Token []OptionalClaim `json:"saml2Token,omitempty"`
}
