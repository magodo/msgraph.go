// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// OMASetting undocumented
type OMASetting struct {
	// Object is the base model of OMASetting
	Object
	// DisplayName Display Name.
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description.
	Description *string `json:"description,omitempty"`
	// OMAURI OMA.
	OMAURI *string `json:"omaUri,omitempty"`
}

// OMASettingBase64 undocumented
type OMASettingBase64 struct {
	// OMASetting is the base model of OMASettingBase64
	OMASetting
	// FileName File name associated with the Value property (*.cer | *.crt | *.p7b | *.bin).
	FileName *string `json:"fileName,omitempty"`
	// Value Value. (Base64 encoded string)
	Value *string `json:"value,omitempty"`
}

// OMASettingBoolean undocumented
type OMASettingBoolean struct {
	// OMASetting is the base model of OMASettingBoolean
	OMASetting
	// Value Value.
	Value *bool `json:"value,omitempty"`
}

// OMASettingDateTime undocumented
type OMASettingDateTime struct {
	// OMASetting is the base model of OMASettingDateTime
	OMASetting
	// Value Value.
	Value *time.Time `json:"value,omitempty"`
}

// OMASettingFloatingPoint undocumented
type OMASettingFloatingPoint struct {
	// OMASetting is the base model of OMASettingFloatingPoint
	OMASetting
	// Value Value.
	Value *float64 `json:"value,omitempty"`
}

// OMASettingInteger undocumented
type OMASettingInteger struct {
	// OMASetting is the base model of OMASettingInteger
	OMASetting
	// Value Value.
	Value *int `json:"value,omitempty"`
}

// OMASettingString undocumented
type OMASettingString struct {
	// OMASetting is the base model of OMASettingString
	OMASetting
	// Value Value.
	Value *string `json:"value,omitempty"`
}

// OMASettingStringXML undocumented
type OMASettingStringXML struct {
	// OMASetting is the base model of OMASettingStringXML
	OMASetting
	// FileName File name associated with the Value property (*.xml).
	FileName *string `json:"fileName,omitempty"`
	// Value Value. (UTF8 encoded byte array)
	Value *Binary `json:"value,omitempty"`
}
