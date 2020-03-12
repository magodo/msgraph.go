// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ProfileRequestBuilder is request builder for Profile
type ProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProfileRequest
func (b *ProfileRequestBuilder) Request() *ProfileRequest {
	return &ProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProfileRequest is request for Profile
type ProfileRequest struct{ BaseRequest }

// Get performs GET request for Profile
func (r *ProfileRequest) Get(ctx context.Context) (resObj *Profile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Profile
func (r *ProfileRequest) Update(ctx context.Context, reqObj *Profile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Profile
func (r *ProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ProfilePhotoRequestBuilder is request builder for ProfilePhoto
type ProfilePhotoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProfilePhotoRequest
func (b *ProfilePhotoRequestBuilder) Request() *ProfilePhotoRequest {
	return &ProfilePhotoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProfilePhotoRequest is request for ProfilePhoto
type ProfilePhotoRequest struct{ BaseRequest }

// Get performs GET request for ProfilePhoto
func (r *ProfilePhotoRequest) Get(ctx context.Context) (resObj *ProfilePhoto, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProfilePhoto
func (r *ProfilePhotoRequest) Update(ctx context.Context, reqObj *ProfilePhoto) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProfilePhoto
func (r *ProfilePhotoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
