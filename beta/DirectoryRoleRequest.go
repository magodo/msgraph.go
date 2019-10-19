// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DirectoryRoleRequestBuilder is request builder for DirectoryRole
type DirectoryRoleRequestBuilder struct{ BaseRequestBuilder }

// Request returns DirectoryRoleRequest
func (b *DirectoryRoleRequestBuilder) Request() *DirectoryRoleRequest {
	return &DirectoryRoleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DirectoryRoleRequest is request for DirectoryRole
type DirectoryRoleRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryRole
func (r *DirectoryRoleRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryRole, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DirectoryRole
func (r *DirectoryRoleRequest) Get() (*DirectoryRole, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DirectoryRole
func (r *DirectoryRoleRequest) Update(reqObj *DirectoryRole) (*DirectoryRole, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DirectoryRole
func (r *DirectoryRoleRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Members returns request builder for DirectoryObject collection
func (b *DirectoryRoleRequestBuilder) Members() *DirectoryRoleMembersCollectionRequestBuilder {
	bb := &DirectoryRoleMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// DirectoryRoleMembersCollectionRequestBuilder is request builder for DirectoryObject collection
type DirectoryRoleMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DirectoryRoleMembersCollectionRequestBuilder) Request() *DirectoryRoleMembersCollectionRequest {
	return &DirectoryRoleMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DirectoryRoleMembersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DirectoryRoleMembersCollectionRequest is request for DirectoryObject collection
type DirectoryRoleMembersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryObject, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DirectoryObject
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Add(reqObj *DirectoryObject) (*DirectoryObject, error) {
	return r.Do("POST", "", reqObj)
}

// ScopedMembers returns request builder for ScopedRoleMembership collection
func (b *DirectoryRoleRequestBuilder) ScopedMembers() *DirectoryRoleScopedMembersCollectionRequestBuilder {
	bb := &DirectoryRoleScopedMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/scopedMembers"
	return bb
}

// DirectoryRoleScopedMembersCollectionRequestBuilder is request builder for ScopedRoleMembership collection
type DirectoryRoleScopedMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ScopedRoleMembership collection
func (b *DirectoryRoleScopedMembersCollectionRequestBuilder) Request() *DirectoryRoleScopedMembersCollectionRequest {
	return &DirectoryRoleScopedMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ScopedRoleMembership item
func (b *DirectoryRoleScopedMembersCollectionRequestBuilder) ID(id string) *ScopedRoleMembershipRequestBuilder {
	bb := &ScopedRoleMembershipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DirectoryRoleScopedMembersCollectionRequest is request for ScopedRoleMembership collection
type DirectoryRoleScopedMembersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ScopedRoleMembership, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Paging(method, path string, obj interface{}) ([]ScopedRoleMembership, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ScopedRoleMembership
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ScopedRoleMembership
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Get() ([]ScopedRoleMembership, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Add(reqObj *ScopedRoleMembership) (*ScopedRoleMembership, error) {
	return r.Do("POST", "", reqObj)
}