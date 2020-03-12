// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OfficeClientConfigurationRequestBuilder is request builder for OfficeClientConfiguration
type OfficeClientConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeClientConfigurationRequest
func (b *OfficeClientConfigurationRequestBuilder) Request() *OfficeClientConfigurationRequest {
	return &OfficeClientConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeClientConfigurationRequest is request for OfficeClientConfiguration
type OfficeClientConfigurationRequest struct{ BaseRequest }

// Get performs GET request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Get(ctx context.Context) (resObj *OfficeClientConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Update(ctx context.Context, reqObj *OfficeClientConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OfficeClientConfigurationAssignmentRequestBuilder is request builder for OfficeClientConfigurationAssignment
type OfficeClientConfigurationAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeClientConfigurationAssignmentRequest
func (b *OfficeClientConfigurationAssignmentRequestBuilder) Request() *OfficeClientConfigurationAssignmentRequest {
	return &OfficeClientConfigurationAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeClientConfigurationAssignmentRequest is request for OfficeClientConfigurationAssignment
type OfficeClientConfigurationAssignmentRequest struct{ BaseRequest }

// Get performs GET request for OfficeClientConfigurationAssignment
func (r *OfficeClientConfigurationAssignmentRequest) Get(ctx context.Context) (resObj *OfficeClientConfigurationAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeClientConfigurationAssignment
func (r *OfficeClientConfigurationAssignmentRequest) Update(ctx context.Context, reqObj *OfficeClientConfigurationAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeClientConfigurationAssignment
func (r *OfficeClientConfigurationAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OfficeConfigurationRequestBuilder is request builder for OfficeConfiguration
type OfficeConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeConfigurationRequest
func (b *OfficeConfigurationRequestBuilder) Request() *OfficeConfigurationRequest {
	return &OfficeConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeConfigurationRequest is request for OfficeConfiguration
type OfficeConfigurationRequest struct{ BaseRequest }

// Get performs GET request for OfficeConfiguration
func (r *OfficeConfigurationRequest) Get(ctx context.Context) (resObj *OfficeConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeConfiguration
func (r *OfficeConfigurationRequest) Update(ctx context.Context, reqObj *OfficeConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeConfiguration
func (r *OfficeConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OfficeGraphInsightsRequestBuilder is request builder for OfficeGraphInsights
type OfficeGraphInsightsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeGraphInsightsRequest
func (b *OfficeGraphInsightsRequestBuilder) Request() *OfficeGraphInsightsRequest {
	return &OfficeGraphInsightsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeGraphInsightsRequest is request for OfficeGraphInsights
type OfficeGraphInsightsRequest struct{ BaseRequest }

// Get performs GET request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Get(ctx context.Context) (resObj *OfficeGraphInsights, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Update(ctx context.Context, reqObj *OfficeGraphInsights) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder struct{ BaseRequestBuilder }

// UpdatePriorities action undocumented
func (b *OfficeConfigurationClientConfigurationsCollectionRequestBuilder) UpdatePriorities(reqObj *OfficeClientConfigurationCollectionUpdatePrioritiesRequestParameter) *OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder {
	bb := &OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updatePriorities"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OfficeClientConfigurationCollectionUpdatePrioritiesRequest struct{ BaseRequest }

//
func (b *OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder) Request() *OfficeClientConfigurationCollectionUpdatePrioritiesRequest {
	return &OfficeClientConfigurationCollectionUpdatePrioritiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OfficeClientConfigurationCollectionUpdatePrioritiesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type OfficeClientConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *OfficeClientConfigurationRequestBuilder) Assign(reqObj *OfficeClientConfigurationAssignRequestParameter) *OfficeClientConfigurationAssignRequestBuilder {
	bb := &OfficeClientConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OfficeClientConfigurationAssignRequest struct{ BaseRequest }

//
func (b *OfficeClientConfigurationAssignRequestBuilder) Request() *OfficeClientConfigurationAssignRequest {
	return &OfficeClientConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OfficeClientConfigurationAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OfficeClientConfigurationAssignment, error) {
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
	var values []OfficeClientConfigurationAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []OfficeClientConfigurationAssignment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
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

//
func (r *OfficeClientConfigurationAssignRequest) Post(ctx context.Context) ([]OfficeClientConfigurationAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}
