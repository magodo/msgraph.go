// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// LicenseAssignmentState undocumented
type LicenseAssignmentState struct {
	// Object is the base model of LicenseAssignmentState
	Object
	// SKUID undocumented
	SKUID *UUID `json:"skuId,omitempty"`
	// DisabledPlans undocumented
	DisabledPlans []UUID `json:"disabledPlans,omitempty"`
	// AssignedByGroup undocumented
	AssignedByGroup *string `json:"assignedByGroup,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// Error undocumented
	Error *string `json:"error,omitempty"`
}

// LicenseDetails undocumented
type LicenseDetails struct {
	// Entity is the base model of LicenseDetails
	Entity
	// ServicePlans undocumented
	ServicePlans []ServicePlanInfo `json:"servicePlans,omitempty"`
	// SKUID undocumented
	SKUID *UUID `json:"skuId,omitempty"`
	// SKUPartNumber undocumented
	SKUPartNumber *string `json:"skuPartNumber,omitempty"`
}

// LicenseProcessingState undocumented
type LicenseProcessingState struct {
	// Object is the base model of LicenseProcessingState
	Object
	// State undocumented
	State *string `json:"state,omitempty"`
}

// LicenseUnitsDetail undocumented
type LicenseUnitsDetail struct {
	// Object is the base model of LicenseUnitsDetail
	Object
	// Enabled undocumented
	Enabled *int `json:"enabled,omitempty"`
	// Suspended undocumented
	Suspended *int `json:"suspended,omitempty"`
	// Warning undocumented
	Warning *int `json:"warning,omitempty"`
}
