// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceKey undocumented
type DeviceKey struct {
	// KeyType undocumented
	KeyType *string `json:"keyType,omitempty"`
	// KeyMaterial undocumented
	KeyMaterial *Binary `json:"keyMaterial,omitempty"`
	// DeviceID undocumented
	DeviceID *UUID `json:"deviceId,omitempty"`
}