// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ServiceHostedMediaConfig undocumented
type ServiceHostedMediaConfig struct {
	// MediaConfig is the base model of ServiceHostedMediaConfig
	MediaConfig
	// PreFetchMedia undocumented
	PreFetchMedia []MediaInfo `json:"preFetchMedia,omitempty"`
}

// ServicePlanInfo undocumented
type ServicePlanInfo struct {
	// Object is the base model of ServicePlanInfo
	Object
	// ServicePlanID undocumented
	ServicePlanID *UUID `json:"servicePlanId,omitempty"`
	// ServicePlanName undocumented
	ServicePlanName *string `json:"servicePlanName,omitempty"`
	// ProvisioningStatus undocumented
	ProvisioningStatus *string `json:"provisioningStatus,omitempty"`
	// AppliesTo undocumented
	AppliesTo *string `json:"appliesTo,omitempty"`
}
