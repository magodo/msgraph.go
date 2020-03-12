// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// RestrictedAppsViolationRequestBuilder is request builder for RestrictedAppsViolation
type RestrictedAppsViolationRequestBuilder struct{ BaseRequestBuilder }

// Request returns RestrictedAppsViolationRequest
func (b *RestrictedAppsViolationRequestBuilder) Request() *RestrictedAppsViolationRequest {
	return &RestrictedAppsViolationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RestrictedAppsViolationRequest is request for RestrictedAppsViolation
type RestrictedAppsViolationRequest struct{ BaseRequest }

// Get performs GET request for RestrictedAppsViolation
func (r *RestrictedAppsViolationRequest) Get(ctx context.Context) (resObj *RestrictedAppsViolation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RestrictedAppsViolation
func (r *RestrictedAppsViolationRequest) Update(ctx context.Context, reqObj *RestrictedAppsViolation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RestrictedAppsViolation
func (r *RestrictedAppsViolationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RestrictedSignInRequestBuilder is request builder for RestrictedSignIn
type RestrictedSignInRequestBuilder struct{ BaseRequestBuilder }

// Request returns RestrictedSignInRequest
func (b *RestrictedSignInRequestBuilder) Request() *RestrictedSignInRequest {
	return &RestrictedSignInRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RestrictedSignInRequest is request for RestrictedSignIn
type RestrictedSignInRequest struct{ BaseRequest }

// Get performs GET request for RestrictedSignIn
func (r *RestrictedSignInRequest) Get(ctx context.Context) (resObj *RestrictedSignIn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RestrictedSignIn
func (r *RestrictedSignInRequest) Update(ctx context.Context, reqObj *RestrictedSignIn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RestrictedSignIn
func (r *RestrictedSignInRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
