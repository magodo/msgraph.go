// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// GovernanceResourceCollectionRegisterRequestParameter undocumented
type GovernanceResourceCollectionRegisterRequestParameter struct {
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
}

// GovernanceRoleAssignmentRequestObjectCancelRequestParameter undocumented
type GovernanceRoleAssignmentRequestObjectCancelRequestParameter struct {
}

// GovernanceRoleAssignmentRequestObjectUpdateRequestActionRequestParameter undocumented
type GovernanceRoleAssignmentRequestObjectUpdateRequestActionRequestParameter struct {
	// Decision undocumented
	Decision *string `json:"decision,omitempty"`
	// AssignmentState undocumented
	AssignmentState *string `json:"assignmentState,omitempty"`
	// Schedule undocumented
	Schedule *GovernanceSchedule `json:"schedule,omitempty"`
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
}

// Parent is navigation property
func (b *GovernanceResourceRequestBuilder) Parent() *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parent"
	return bb
}

// RoleAssignmentRequests returns request builder for GovernanceRoleAssignmentRequestObject collection
func (b *GovernanceResourceRequestBuilder) RoleAssignmentRequests() *GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder {
	bb := &GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignmentRequests"
	return bb
}

// GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder is request builder for GovernanceRoleAssignmentRequestObject collection
type GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleAssignmentRequestObject collection
func (b *GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder) Request() *GovernanceResourceRoleAssignmentRequestsCollectionRequest {
	return &GovernanceResourceRoleAssignmentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleAssignmentRequestObject item
func (b *GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder) ID(id string) *GovernanceRoleAssignmentRequestObjectRequestBuilder {
	bb := &GovernanceRoleAssignmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GovernanceResourceRoleAssignmentRequestsCollectionRequest is request for GovernanceRoleAssignmentRequestObject collection
type GovernanceResourceRoleAssignmentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleAssignmentRequestObject collection
func (r *GovernanceResourceRoleAssignmentRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernanceRoleAssignmentRequestObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleAssignmentRequestObject
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []GovernanceRoleAssignmentRequestObject
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for GovernanceRoleAssignmentRequestObject collection, max N pages
func (r *GovernanceResourceRoleAssignmentRequestsCollectionRequest) GetN(ctx context.Context, n int) ([]GovernanceRoleAssignmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernanceRoleAssignmentRequestObject collection
func (r *GovernanceResourceRoleAssignmentRequestsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleAssignmentRequestObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernanceRoleAssignmentRequestObject collection
func (r *GovernanceResourceRoleAssignmentRequestsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleAssignmentRequestObject) (resObj *GovernanceRoleAssignmentRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleAssignments returns request builder for GovernanceRoleAssignment collection
func (b *GovernanceResourceRequestBuilder) RoleAssignments() *GovernanceResourceRoleAssignmentsCollectionRequestBuilder {
	bb := &GovernanceResourceRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignments"
	return bb
}

// GovernanceResourceRoleAssignmentsCollectionRequestBuilder is request builder for GovernanceRoleAssignment collection
type GovernanceResourceRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleAssignment collection
func (b *GovernanceResourceRoleAssignmentsCollectionRequestBuilder) Request() *GovernanceResourceRoleAssignmentsCollectionRequest {
	return &GovernanceResourceRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleAssignment item
func (b *GovernanceResourceRoleAssignmentsCollectionRequestBuilder) ID(id string) *GovernanceRoleAssignmentRequestBuilder {
	bb := &GovernanceRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GovernanceResourceRoleAssignmentsCollectionRequest is request for GovernanceRoleAssignment collection
type GovernanceResourceRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleAssignment collection
func (r *GovernanceResourceRoleAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernanceRoleAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleAssignment
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []GovernanceRoleAssignment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for GovernanceRoleAssignment collection, max N pages
func (r *GovernanceResourceRoleAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]GovernanceRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernanceRoleAssignment collection
func (r *GovernanceResourceRoleAssignmentsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernanceRoleAssignment collection
func (r *GovernanceResourceRoleAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleAssignment) (resObj *GovernanceRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleDefinitions returns request builder for GovernanceRoleDefinition collection
func (b *GovernanceResourceRequestBuilder) RoleDefinitions() *GovernanceResourceRoleDefinitionsCollectionRequestBuilder {
	bb := &GovernanceResourceRoleDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinitions"
	return bb
}

// GovernanceResourceRoleDefinitionsCollectionRequestBuilder is request builder for GovernanceRoleDefinition collection
type GovernanceResourceRoleDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleDefinition collection
func (b *GovernanceResourceRoleDefinitionsCollectionRequestBuilder) Request() *GovernanceResourceRoleDefinitionsCollectionRequest {
	return &GovernanceResourceRoleDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleDefinition item
func (b *GovernanceResourceRoleDefinitionsCollectionRequestBuilder) ID(id string) *GovernanceRoleDefinitionRequestBuilder {
	bb := &GovernanceRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GovernanceResourceRoleDefinitionsCollectionRequest is request for GovernanceRoleDefinition collection
type GovernanceResourceRoleDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleDefinition collection
func (r *GovernanceResourceRoleDefinitionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernanceRoleDefinition, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleDefinition
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []GovernanceRoleDefinition
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for GovernanceRoleDefinition collection, max N pages
func (r *GovernanceResourceRoleDefinitionsCollectionRequest) GetN(ctx context.Context, n int) ([]GovernanceRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernanceRoleDefinition collection
func (r *GovernanceResourceRoleDefinitionsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernanceRoleDefinition collection
func (r *GovernanceResourceRoleDefinitionsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleDefinition) (resObj *GovernanceRoleDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleSettings returns request builder for GovernanceRoleSetting collection
func (b *GovernanceResourceRequestBuilder) RoleSettings() *GovernanceResourceRoleSettingsCollectionRequestBuilder {
	bb := &GovernanceResourceRoleSettingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleSettings"
	return bb
}

// GovernanceResourceRoleSettingsCollectionRequestBuilder is request builder for GovernanceRoleSetting collection
type GovernanceResourceRoleSettingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleSetting collection
func (b *GovernanceResourceRoleSettingsCollectionRequestBuilder) Request() *GovernanceResourceRoleSettingsCollectionRequest {
	return &GovernanceResourceRoleSettingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleSetting item
func (b *GovernanceResourceRoleSettingsCollectionRequestBuilder) ID(id string) *GovernanceRoleSettingRequestBuilder {
	bb := &GovernanceRoleSettingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GovernanceResourceRoleSettingsCollectionRequest is request for GovernanceRoleSetting collection
type GovernanceResourceRoleSettingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleSetting collection
func (r *GovernanceResourceRoleSettingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernanceRoleSetting, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleSetting
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []GovernanceRoleSetting
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for GovernanceRoleSetting collection, max N pages
func (r *GovernanceResourceRoleSettingsCollectionRequest) GetN(ctx context.Context, n int) ([]GovernanceRoleSetting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernanceRoleSetting collection
func (r *GovernanceResourceRoleSettingsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleSetting, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernanceRoleSetting collection
func (r *GovernanceResourceRoleSettingsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleSetting) (resObj *GovernanceRoleSetting, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// LinkedEligibleRoleAssignment is navigation property
func (b *GovernanceRoleAssignmentRequestBuilder) LinkedEligibleRoleAssignment() *GovernanceRoleAssignmentRequestBuilder {
	bb := &GovernanceRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/linkedEligibleRoleAssignment"
	return bb
}

// Resource is navigation property
func (b *GovernanceRoleAssignmentRequestBuilder) Resource() *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}

// RoleDefinition is navigation property
func (b *GovernanceRoleAssignmentRequestBuilder) RoleDefinition() *GovernanceRoleDefinitionRequestBuilder {
	bb := &GovernanceRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// Subject is navigation property
func (b *GovernanceRoleAssignmentRequestBuilder) Subject() *GovernanceSubjectRequestBuilder {
	bb := &GovernanceSubjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/subject"
	return bb
}

// Resource is navigation property
func (b *GovernanceRoleAssignmentRequestObjectRequestBuilder) Resource() *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}

// RoleDefinition is navigation property
func (b *GovernanceRoleAssignmentRequestObjectRequestBuilder) RoleDefinition() *GovernanceRoleDefinitionRequestBuilder {
	bb := &GovernanceRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// Subject is navigation property
func (b *GovernanceRoleAssignmentRequestObjectRequestBuilder) Subject() *GovernanceSubjectRequestBuilder {
	bb := &GovernanceSubjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/subject"
	return bb
}

// Resource is navigation property
func (b *GovernanceRoleDefinitionRequestBuilder) Resource() *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}

// RoleSetting is navigation property
func (b *GovernanceRoleDefinitionRequestBuilder) RoleSetting() *GovernanceRoleSettingRequestBuilder {
	bb := &GovernanceRoleSettingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleSetting"
	return bb
}

// Resource is navigation property
func (b *GovernanceRoleSettingRequestBuilder) Resource() *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}

// RoleDefinition is navigation property
func (b *GovernanceRoleSettingRequestBuilder) RoleDefinition() *GovernanceRoleDefinitionRequestBuilder {
	bb := &GovernanceRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}
