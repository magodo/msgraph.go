// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// OnPremisesConditionalAccessSettings Singleton entity which represents the Exchange OnPremises Conditional Access Settings for a tenant.
type OnPremisesConditionalAccessSettings struct {
	// Entity is the base model of OnPremisesConditionalAccessSettings
	Entity
	// Enabled Indicates if on premises conditional access is enabled for this organization
	Enabled *bool `json:"enabled,omitempty"`
	// IncludedGroups User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access.
	IncludedGroups []UUID `json:"includedGroups,omitempty"`
	// ExcludedGroups User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy.
	ExcludedGroups []UUID `json:"excludedGroups,omitempty"`
	// OverrideDefaultRule Override the default access rule when allowing a device to ensure access is granted.
	OverrideDefaultRule *bool `json:"overrideDefaultRule,omitempty"`
}

// OnPremisesExtensionAttributes undocumented
type OnPremisesExtensionAttributes struct {
	// Object is the base model of OnPremisesExtensionAttributes
	Object
	// ExtensionAttribute1 undocumented
	ExtensionAttribute1 *string `json:"extensionAttribute1,omitempty"`
	// ExtensionAttribute2 undocumented
	ExtensionAttribute2 *string `json:"extensionAttribute2,omitempty"`
	// ExtensionAttribute3 undocumented
	ExtensionAttribute3 *string `json:"extensionAttribute3,omitempty"`
	// ExtensionAttribute4 undocumented
	ExtensionAttribute4 *string `json:"extensionAttribute4,omitempty"`
	// ExtensionAttribute5 undocumented
	ExtensionAttribute5 *string `json:"extensionAttribute5,omitempty"`
	// ExtensionAttribute6 undocumented
	ExtensionAttribute6 *string `json:"extensionAttribute6,omitempty"`
	// ExtensionAttribute7 undocumented
	ExtensionAttribute7 *string `json:"extensionAttribute7,omitempty"`
	// ExtensionAttribute8 undocumented
	ExtensionAttribute8 *string `json:"extensionAttribute8,omitempty"`
	// ExtensionAttribute9 undocumented
	ExtensionAttribute9 *string `json:"extensionAttribute9,omitempty"`
	// ExtensionAttribute10 undocumented
	ExtensionAttribute10 *string `json:"extensionAttribute10,omitempty"`
	// ExtensionAttribute11 undocumented
	ExtensionAttribute11 *string `json:"extensionAttribute11,omitempty"`
	// ExtensionAttribute12 undocumented
	ExtensionAttribute12 *string `json:"extensionAttribute12,omitempty"`
	// ExtensionAttribute13 undocumented
	ExtensionAttribute13 *string `json:"extensionAttribute13,omitempty"`
	// ExtensionAttribute14 undocumented
	ExtensionAttribute14 *string `json:"extensionAttribute14,omitempty"`
	// ExtensionAttribute15 undocumented
	ExtensionAttribute15 *string `json:"extensionAttribute15,omitempty"`
}

// OnPremisesProvisioningError undocumented
type OnPremisesProvisioningError struct {
	// Object is the base model of OnPremisesProvisioningError
	Object
	// Value undocumented
	Value *string `json:"value,omitempty"`
	// Category undocumented
	Category *string `json:"category,omitempty"`
	// PropertyCausingError undocumented
	PropertyCausingError *string `json:"propertyCausingError,omitempty"`
	// OccurredDateTime undocumented
	OccurredDateTime *time.Time `json:"occurredDateTime,omitempty"`
}
