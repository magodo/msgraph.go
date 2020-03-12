// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ConditionalAccessPolicyRequestBuilder is request builder for ConditionalAccessPolicy
type ConditionalAccessPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessPolicyRequest
func (b *ConditionalAccessPolicyRequestBuilder) Request() *ConditionalAccessPolicyRequest {
	return &ConditionalAccessPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessPolicyRequest is request for ConditionalAccessPolicy
type ConditionalAccessPolicyRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessPolicy
func (r *ConditionalAccessPolicyRequest) Get(ctx context.Context) (resObj *ConditionalAccessPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessPolicy
func (r *ConditionalAccessPolicyRequest) Update(ctx context.Context, reqObj *ConditionalAccessPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessPolicy
func (r *ConditionalAccessPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessRootRequestBuilder is request builder for ConditionalAccessRoot
type ConditionalAccessRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessRootRequest
func (b *ConditionalAccessRootRequestBuilder) Request() *ConditionalAccessRootRequest {
	return &ConditionalAccessRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessRootRequest is request for ConditionalAccessRoot
type ConditionalAccessRootRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessRoot
func (r *ConditionalAccessRootRequest) Get(ctx context.Context) (resObj *ConditionalAccessRoot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessRoot
func (r *ConditionalAccessRootRequest) Update(ctx context.Context, reqObj *ConditionalAccessRoot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessRoot
func (r *ConditionalAccessRootRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
