// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SensitivityLabelRequestBuilder is request builder for SensitivityLabel
type SensitivityLabelRequestBuilder struct{ BaseRequestBuilder }

// Request returns SensitivityLabelRequest
func (b *SensitivityLabelRequestBuilder) Request() *SensitivityLabelRequest {
	return &SensitivityLabelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SensitivityLabelRequest is request for SensitivityLabel
type SensitivityLabelRequest struct{ BaseRequest }

// Get performs GET request for SensitivityLabel
func (r *SensitivityLabelRequest) Get(ctx context.Context) (resObj *SensitivityLabel, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SensitivityLabel
func (r *SensitivityLabelRequest) Update(ctx context.Context, reqObj *SensitivityLabel) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SensitivityLabel
func (r *SensitivityLabelRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SensitivityPolicySettingsRequestBuilder is request builder for SensitivityPolicySettings
type SensitivityPolicySettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns SensitivityPolicySettingsRequest
func (b *SensitivityPolicySettingsRequestBuilder) Request() *SensitivityPolicySettingsRequest {
	return &SensitivityPolicySettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SensitivityPolicySettingsRequest is request for SensitivityPolicySettings
type SensitivityPolicySettingsRequest struct{ BaseRequest }

// Get performs GET request for SensitivityPolicySettings
func (r *SensitivityPolicySettingsRequest) Get(ctx context.Context) (resObj *SensitivityPolicySettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SensitivityPolicySettings
func (r *SensitivityPolicySettingsRequest) Update(ctx context.Context, reqObj *SensitivityPolicySettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SensitivityPolicySettings
func (r *SensitivityPolicySettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type SensitivityLabelCollectionEvaluateRequestBuilder struct{ BaseRequestBuilder }

// Evaluate action undocumented
func (b *DataClassificationServiceSensitivityLabelsCollectionRequestBuilder) Evaluate(reqObj *SensitivityLabelCollectionEvaluateRequestParameter) *SensitivityLabelCollectionEvaluateRequestBuilder {
	bb := &SensitivityLabelCollectionEvaluateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Evaluate action undocumented
func (b *InformationProtectionSensitivityLabelsCollectionRequestBuilder) Evaluate(reqObj *SensitivityLabelCollectionEvaluateRequestParameter) *SensitivityLabelCollectionEvaluateRequestBuilder {
	bb := &SensitivityLabelCollectionEvaluateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Evaluate action undocumented
func (b *SensitivityLabelSublabelsCollectionRequestBuilder) Evaluate(reqObj *SensitivityLabelCollectionEvaluateRequestParameter) *SensitivityLabelCollectionEvaluateRequestBuilder {
	bb := &SensitivityLabelCollectionEvaluateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SensitivityLabelCollectionEvaluateRequest struct{ BaseRequest }

//
func (b *SensitivityLabelCollectionEvaluateRequestBuilder) Request() *SensitivityLabelCollectionEvaluateRequest {
	return &SensitivityLabelCollectionEvaluateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SensitivityLabelCollectionEvaluateRequest) Post(ctx context.Context) (resObj *EvaluateLabelJobResponse, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
