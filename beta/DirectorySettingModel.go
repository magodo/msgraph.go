// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DirectorySetting undocumented
type DirectorySetting struct {
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// Values undocumented
	Values []SettingValue `json:"values,omitempty"`
}