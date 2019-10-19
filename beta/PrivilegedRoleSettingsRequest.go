// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PrivilegedRoleSettingsRequestBuilder is request builder for PrivilegedRoleSettings
type PrivilegedRoleSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleSettingsRequest
func (b *PrivilegedRoleSettingsRequestBuilder) Request() *PrivilegedRoleSettingsRequest {
	return &PrivilegedRoleSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedRoleSettingsRequest is request for PrivilegedRoleSettings
type PrivilegedRoleSettingsRequest struct{ BaseRequest }

// Do performs HTTP request for PrivilegedRoleSettings
func (r *PrivilegedRoleSettingsRequest) Do(method, path string, reqObj interface{}) (resObj *PrivilegedRoleSettings, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for PrivilegedRoleSettings
func (r *PrivilegedRoleSettingsRequest) Get() (*PrivilegedRoleSettings, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for PrivilegedRoleSettings
func (r *PrivilegedRoleSettingsRequest) Update(reqObj *PrivilegedRoleSettings) (*PrivilegedRoleSettings, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for PrivilegedRoleSettings
func (r *PrivilegedRoleSettingsRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}