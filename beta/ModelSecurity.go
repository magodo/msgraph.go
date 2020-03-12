// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Security undocumented
type Security struct {
	// Entity is the base model of Security
	Entity
	// ProviderStatus undocumented
	ProviderStatus []SecurityProviderStatus `json:"providerStatus,omitempty"`
	// Alerts undocumented
	Alerts []Alert `json:"alerts,omitempty"`
	// CloudAppSecurityProfiles undocumented
	CloudAppSecurityProfiles []CloudAppSecurityProfile `json:"cloudAppSecurityProfiles,omitempty"`
	// DomainSecurityProfiles undocumented
	DomainSecurityProfiles []DomainSecurityProfile `json:"domainSecurityProfiles,omitempty"`
	// FileSecurityProfiles undocumented
	FileSecurityProfiles []FileSecurityProfile `json:"fileSecurityProfiles,omitempty"`
	// HostSecurityProfiles undocumented
	HostSecurityProfiles []HostSecurityProfile `json:"hostSecurityProfiles,omitempty"`
	// IPSecurityProfiles undocumented
	IPSecurityProfiles []IPSecurityProfile `json:"ipSecurityProfiles,omitempty"`
	// ProviderTenantSettings undocumented
	ProviderTenantSettings []ProviderTenantSetting `json:"providerTenantSettings,omitempty"`
	// SecureScoreControlProfiles undocumented
	SecureScoreControlProfiles []SecureScoreControlProfile `json:"secureScoreControlProfiles,omitempty"`
	// SecureScores undocumented
	SecureScores []SecureScore `json:"secureScores,omitempty"`
	// TiIndicators undocumented
	TiIndicators []TiIndicator `json:"tiIndicators,omitempty"`
	// UserSecurityProfiles undocumented
	UserSecurityProfiles []UserSecurityProfile `json:"userSecurityProfiles,omitempty"`
	// SecurityActions undocumented
	SecurityActions []SecurityAction `json:"securityActions,omitempty"`
}

// SecurityAction undocumented
type SecurityAction struct {
	// Entity is the base model of SecurityAction
	Entity
	// ActionReason undocumented
	ActionReason *string `json:"actionReason,omitempty"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
	// CompletedDateTime undocumented
	CompletedDateTime *time.Time `json:"completedDateTime,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ErrorInfo undocumented
	ErrorInfo *ResultInfo `json:"errorInfo,omitempty"`
	// LastActionDateTime undocumented
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Parameters undocumented
	Parameters []KeyValuePair `json:"parameters,omitempty"`
	// States undocumented
	States []SecurityActionState `json:"states,omitempty"`
	// Status undocumented
	Status *OperationStatus `json:"status,omitempty"`
	// User undocumented
	User *string `json:"user,omitempty"`
	// VendorInformation undocumented
	VendorInformation *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}

// SecurityActionState undocumented
type SecurityActionState struct {
	// Object is the base model of SecurityActionState
	Object
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// Status undocumented
	Status *OperationStatus `json:"status,omitempty"`
	// UpdatedDateTime undocumented
	UpdatedDateTime *time.Time `json:"updatedDateTime,omitempty"`
	// User undocumented
	User *string `json:"user,omitempty"`
}

// SecurityBaselineCategoryStateSummary The security baseline per category compliance state summary for the security baseline of the account.
type SecurityBaselineCategoryStateSummary struct {
	// SecurityBaselineStateSummary is the base model of SecurityBaselineCategoryStateSummary
	SecurityBaselineStateSummary
	// DisplayName The category name
	DisplayName *string `json:"displayName,omitempty"`
}

// SecurityBaselineDeviceState The security baseline compliance state summary of the security baseline for a device.
type SecurityBaselineDeviceState struct {
	// Entity is the base model of SecurityBaselineDeviceState
	Entity
	// ManagedDeviceID Intune device id
	ManagedDeviceID *string `json:"managedDeviceId,omitempty"`
	// DeviceDisplayName Display name of the device
	DeviceDisplayName *string `json:"deviceDisplayName,omitempty"`
	// UserPrincipalName User Principal Name
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// State Security baseline compliance state
	State *SecurityBaselineComplianceState `json:"state,omitempty"`
	// LastReportedDateTime Last modified date time of the policy report
	LastReportedDateTime *time.Time `json:"lastReportedDateTime,omitempty"`
}

// SecurityBaselineSettingState The security baseline compliance state of a setting for a device
type SecurityBaselineSettingState struct {
	// Entity is the base model of SecurityBaselineSettingState
	Entity
	// SettingName The setting name that is being reported
	SettingName *string `json:"settingName,omitempty"`
	// State The compliance state of the security baseline setting
	State *SecurityBaselineComplianceState `json:"state,omitempty"`
	// SettingCategoryID The setting category id which this setting belongs to
	SettingCategoryID *string `json:"settingCategoryId,omitempty"`
}

// SecurityBaselineState Security baseline state for a device.
type SecurityBaselineState struct {
	// Entity is the base model of SecurityBaselineState
	Entity
	// SecurityBaselineTemplateID The security baseline template id
	SecurityBaselineTemplateID *string `json:"securityBaselineTemplateId,omitempty"`
	// DisplayName The display name of the security baseline
	DisplayName *string `json:"displayName,omitempty"`
	// SettingStates undocumented
	SettingStates []SecurityBaselineSettingState `json:"settingStates,omitempty"`
}

// SecurityBaselineStateSummary The security baseline compliance state summary for the security baseline of the account.
type SecurityBaselineStateSummary struct {
	// Entity is the base model of SecurityBaselineStateSummary
	Entity
	// SecureCount Number of secure devices
	SecureCount *int `json:"secureCount,omitempty"`
	// NotSecureCount Number of not secure devices
	NotSecureCount *int `json:"notSecureCount,omitempty"`
	// UnknownCount Number of unknown devices
	UnknownCount *int `json:"unknownCount,omitempty"`
	// ErrorCount Number of error devices
	ErrorCount *int `json:"errorCount,omitempty"`
	// ConflictCount Number of conflict devices
	ConflictCount *int `json:"conflictCount,omitempty"`
	// NotApplicableCount Number of not applicable devices
	NotApplicableCount *int `json:"notApplicableCount,omitempty"`
}

// SecurityBaselineTemplate The security baseline template of the account
type SecurityBaselineTemplate struct {
	// DeviceManagementTemplate is the base model of SecurityBaselineTemplate
	DeviceManagementTemplate
	// DeviceStateSummary undocumented
	DeviceStateSummary *SecurityBaselineStateSummary `json:"deviceStateSummary,omitempty"`
	// DeviceStates undocumented
	DeviceStates []SecurityBaselineDeviceState `json:"deviceStates,omitempty"`
	// CategoryDeviceStateSummaries undocumented
	CategoryDeviceStateSummaries []SecurityBaselineCategoryStateSummary `json:"categoryDeviceStateSummaries,omitempty"`
}

// SecurityProviderStatus undocumented
type SecurityProviderStatus struct {
	// Object is the base model of SecurityProviderStatus
	Object
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// Endpoint undocumented
	Endpoint *string `json:"endpoint,omitempty"`
	// Provider undocumented
	Provider *string `json:"provider,omitempty"`
	// Region undocumented
	Region *string `json:"region,omitempty"`
	// Vendor undocumented
	Vendor *string `json:"vendor,omitempty"`
}

// SecurityVendorInformation undocumented
type SecurityVendorInformation struct {
	// Object is the base model of SecurityVendorInformation
	Object
	// Provider undocumented
	Provider *string `json:"provider,omitempty"`
	// ProviderVersion undocumented
	ProviderVersion *string `json:"providerVersion,omitempty"`
	// SubProvider undocumented
	SubProvider *string `json:"subProvider,omitempty"`
	// Vendor undocumented
	Vendor *string `json:"vendor,omitempty"`
}
