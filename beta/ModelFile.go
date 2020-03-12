// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// File undocumented
type File struct {
	// Object is the base model of File
	Object
	// Hashes undocumented
	Hashes *Hashes `json:"hashes,omitempty"`
	// MimeType undocumented
	MimeType *string `json:"mimeType,omitempty"`
	// ProcessingMetadata undocumented
	ProcessingMetadata *bool `json:"processingMetadata,omitempty"`
}

// FileAssessmentRequestObject undocumented
type FileAssessmentRequestObject struct {
	// ThreatAssessmentRequestObject is the base model of FileAssessmentRequestObject
	ThreatAssessmentRequestObject
	// FileName undocumented
	FileName *string `json:"fileName,omitempty"`
	// ContentData undocumented
	ContentData *string `json:"contentData,omitempty"`
}

// FileAttachment undocumented
type FileAttachment struct {
	// Attachment is the base model of FileAttachment
	Attachment
	// ContentID undocumented
	ContentID *string `json:"contentId,omitempty"`
	// ContentLocation undocumented
	ContentLocation *string `json:"contentLocation,omitempty"`
	// ContentBytes undocumented
	ContentBytes *Binary `json:"contentBytes,omitempty"`
}

// FileClassificationRequestObject undocumented
type FileClassificationRequestObject struct {
	// Entity is the base model of FileClassificationRequestObject
	Entity
	// File undocumented
	File *Stream `json:"file,omitempty"`
	// SensitiveTypeIDs undocumented
	SensitiveTypeIDs []string `json:"sensitiveTypeIds,omitempty"`
}

// FileEncryptionInfo undocumented
type FileEncryptionInfo struct {
	// Object is the base model of FileEncryptionInfo
	Object
	// EncryptionKey The key used to encrypt the file content.
	EncryptionKey *Binary `json:"encryptionKey,omitempty"`
	// InitializationVector The initialization vector used for the encryption algorithm.
	InitializationVector *Binary `json:"initializationVector,omitempty"`
	// Mac The hash of the encrypted file content + IV (content hash).
	Mac *Binary `json:"mac,omitempty"`
	// MacKey The key used to get mac.
	MacKey *Binary `json:"macKey,omitempty"`
	// ProfileIdentifier The the profile identifier.
	ProfileIdentifier *string `json:"profileIdentifier,omitempty"`
	// FileDigest The file digest prior to encryption.
	FileDigest *Binary `json:"fileDigest,omitempty"`
	// FileDigestAlgorithm The file digest algorithm.
	FileDigestAlgorithm *string `json:"fileDigestAlgorithm,omitempty"`
}

// FileHash undocumented
type FileHash struct {
	// Object is the base model of FileHash
	Object
	// HashType undocumented
	HashType *FileHashType `json:"hashType,omitempty"`
	// HashValue undocumented
	HashValue *string `json:"hashValue,omitempty"`
}

// FileSecurityProfile undocumented
type FileSecurityProfile struct {
	// Entity is the base model of FileSecurityProfile
	Entity
	// ActivityGroupNames undocumented
	ActivityGroupNames []string `json:"activityGroupNames,omitempty"`
	// AzureSubscriptionID undocumented
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty"`
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// CertificateThumbprint undocumented
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
	// Extensions undocumented
	Extensions []string `json:"extensions,omitempty"`
	// FileType undocumented
	FileType *string `json:"fileType,omitempty"`
	// FirstSeenDateTime undocumented
	FirstSeenDateTime *time.Time `json:"firstSeenDateTime,omitempty"`
	// Hashes undocumented
	Hashes []FileHash `json:"hashes,omitempty"`
	// LastSeenDateTime undocumented
	LastSeenDateTime *time.Time `json:"lastSeenDateTime,omitempty"`
	// MalwareStates undocumented
	MalwareStates []MalwareState `json:"malwareStates,omitempty"`
	// Names undocumented
	Names []string `json:"names,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// VendorInformation undocumented
	VendorInformation *SecurityVendorInformation `json:"vendorInformation,omitempty"`
	// VulnerabilityStates undocumented
	VulnerabilityStates []VulnerabilityState `json:"vulnerabilityStates,omitempty"`
}

// FileSecurityState undocumented
type FileSecurityState struct {
	// Object is the base model of FileSecurityState
	Object
	// FileHash undocumented
	FileHash *FileHash `json:"fileHash,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Path undocumented
	Path *string `json:"path,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
}

// FileSystemInfo undocumented
type FileSystemInfo struct {
	// Object is the base model of FileSystemInfo
	Object
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastAccessedDateTime undocumented
	LastAccessedDateTime *time.Time `json:"lastAccessedDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}
