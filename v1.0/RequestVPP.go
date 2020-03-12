// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// VPPTokenRequestBuilder is request builder for VPPToken
type VPPTokenRequestBuilder struct{ BaseRequestBuilder }

// Request returns VPPTokenRequest
func (b *VPPTokenRequestBuilder) Request() *VPPTokenRequest {
	return &VPPTokenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// VPPTokenRequest is request for VPPToken
type VPPTokenRequest struct{ BaseRequest }

// Get performs GET request for VPPToken
func (r *VPPTokenRequest) Get(ctx context.Context) (resObj *VPPToken, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for VPPToken
func (r *VPPTokenRequest) Update(ctx context.Context, reqObj *VPPToken) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for VPPToken
func (r *VPPTokenRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type VPPTokenSyncLicensesRequestBuilder struct{ BaseRequestBuilder }

// SyncLicenses action undocumented
func (b *VPPTokenRequestBuilder) SyncLicenses(reqObj *VPPTokenSyncLicensesRequestParameter) *VPPTokenSyncLicensesRequestBuilder {
	bb := &VPPTokenSyncLicensesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/syncLicenses"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type VPPTokenSyncLicensesRequest struct{ BaseRequest }

//
func (b *VPPTokenSyncLicensesRequestBuilder) Request() *VPPTokenSyncLicensesRequest {
	return &VPPTokenSyncLicensesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *VPPTokenSyncLicensesRequest) Post(ctx context.Context) (resObj *VPPToken, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
