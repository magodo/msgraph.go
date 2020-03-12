// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ClassificationError undocumented
type ClassificationError struct {
	// ClassifcationErrorBase is the base model of ClassificationError
	ClassifcationErrorBase
	// Details undocumented
	Details []ClassifcationErrorBase `json:"details,omitempty"`
}

// ClassificationInnerError undocumented
type ClassificationInnerError struct {
	// Object is the base model of ClassificationInnerError
	Object
	// ErrorDateTime undocumented
	ErrorDateTime *time.Time `json:"errorDateTime,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// ClientRequestID undocumented
	ClientRequestID *string `json:"clientRequestId,omitempty"`
	// ActivityID undocumented
	ActivityID *string `json:"activityId,omitempty"`
}

// ClassificationJobResponse undocumented
type ClassificationJobResponse struct {
	// JobResponseBase is the base model of ClassificationJobResponse
	JobResponseBase
	// Result undocumented
	Result *DetectedSensitiveContentWrapper `json:"result,omitempty"`
}

// ClassificationResult undocumented
type ClassificationResult struct {
	// Object is the base model of ClassificationResult
	Object
	// SensitiveTypeID undocumented
	SensitiveTypeID *UUID `json:"sensitiveTypeId,omitempty"`
	// Count undocumented
	Count *int `json:"count,omitempty"`
	// ConfidenceLevel undocumented
	ConfidenceLevel *int `json:"confidenceLevel,omitempty"`
}
