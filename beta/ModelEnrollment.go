// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EnrollmentConfigurationAssignment Enrollment Configuration Assignment
type EnrollmentConfigurationAssignment struct {
	// Entity is the base model of EnrollmentConfigurationAssignment
	Entity
	// Target Represents an assignment to managed devices in the tenant
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
	// Source Type of resource used for deployment to a group, direct or policySet
	Source *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`
	// SourceID Identifier for resource used for deployment to a group
	SourceID *string `json:"sourceId,omitempty"`
}

// EnrollmentProfile The enrollmentProfile resource represents a collection of configurations which must be provided pre-enrollment to enable enrolling certain devices whose identities have been pre-staged. Pre-staged device identities are assigned to this type of profile to apply the profile's configurations at enrollment of the corresponding device.
type EnrollmentProfile struct {
	// Entity is the base model of EnrollmentProfile
	Entity
	// DisplayName Name of the profile
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of the profile
	Description *string `json:"description,omitempty"`
	// RequiresUserAuthentication Indicates if the profile requires user authentication
	RequiresUserAuthentication *bool `json:"requiresUserAuthentication,omitempty"`
	// ConfigurationEndpointURL Configuration endpoint url to use for Enrollment
	ConfigurationEndpointURL *string `json:"configurationEndpointUrl,omitempty"`
	// EnableAuthenticationViaCompanyPortal Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
	EnableAuthenticationViaCompanyPortal *bool `json:"enableAuthenticationViaCompanyPortal,omitempty"`
	// RequireCompanyPortalOnSetupAssistantEnrolledDevices Indicates that Company Portal is required on setup assistant enrolled devices
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`
}

// EnrollmentRestrictionsConfigurationPolicySetItem A class containing the properties used for enrollment restriction PolicySetItem.
type EnrollmentRestrictionsConfigurationPolicySetItem struct {
	// PolicySetItem is the base model of EnrollmentRestrictionsConfigurationPolicySetItem
	PolicySetItem
	// Priority Priority of the EnrollmentRestrictionsConfigurationPolicySetItem.
	Priority *int `json:"priority,omitempty"`
	// Limit Limit of the EnrollmentRestrictionsConfigurationPolicySetItem.
	Limit *int `json:"limit,omitempty"`
}

// EnrollmentTroubleshootingEvent Event representing an enrollment failure.
type EnrollmentTroubleshootingEvent struct {
	// DeviceManagementTroubleshootingEvent is the base model of EnrollmentTroubleshootingEvent
	DeviceManagementTroubleshootingEvent
	// ManagedDeviceIdentifier Device identifier created or collected by Intune.
	ManagedDeviceIdentifier *string `json:"managedDeviceIdentifier,omitempty"`
	// OperatingSystem Operating System.
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// OsVersion OS Version.
	OsVersion *string `json:"osVersion,omitempty"`
	// UserID Identifier for the user that tried to enroll the device.
	UserID *string `json:"userId,omitempty"`
	// DeviceID Azure AD device identifier.
	DeviceID *string `json:"deviceId,omitempty"`
	// EnrollmentType Type of the enrollment.
	EnrollmentType *DeviceEnrollmentType `json:"enrollmentType,omitempty"`
	// FailureCategory Highlevel failure category.
	FailureCategory *DeviceEnrollmentFailureReason `json:"failureCategory,omitempty"`
	// FailureReason Detailed failure reason.
	FailureReason *string `json:"failureReason,omitempty"`
}
