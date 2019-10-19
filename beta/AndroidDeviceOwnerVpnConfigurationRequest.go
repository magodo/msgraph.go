// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidDeviceOwnerVpnConfigurationRequestBuilder is request builder for AndroidDeviceOwnerVpnConfiguration
type AndroidDeviceOwnerVpnConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidDeviceOwnerVpnConfigurationRequest
func (b *AndroidDeviceOwnerVpnConfigurationRequestBuilder) Request() *AndroidDeviceOwnerVpnConfigurationRequest {
	return &AndroidDeviceOwnerVpnConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidDeviceOwnerVpnConfigurationRequest is request for AndroidDeviceOwnerVpnConfiguration
type AndroidDeviceOwnerVpnConfigurationRequest struct{ BaseRequest }

// Do performs HTTP request for AndroidDeviceOwnerVpnConfiguration
func (r *AndroidDeviceOwnerVpnConfigurationRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidDeviceOwnerVpnConfiguration, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidDeviceOwnerVpnConfiguration
func (r *AndroidDeviceOwnerVpnConfigurationRequest) Get() (*AndroidDeviceOwnerVpnConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidDeviceOwnerVpnConfiguration
func (r *AndroidDeviceOwnerVpnConfigurationRequest) Update(reqObj *AndroidDeviceOwnerVpnConfiguration) (*AndroidDeviceOwnerVpnConfiguration, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidDeviceOwnerVpnConfiguration
func (r *AndroidDeviceOwnerVpnConfigurationRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// IdentityCertificate is navigation property
func (b *AndroidDeviceOwnerVpnConfigurationRequestBuilder) IdentityCertificate() *AndroidDeviceOwnerCertificateProfileBaseRequestBuilder {
	bb := &AndroidDeviceOwnerCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificate"
	return bb
}