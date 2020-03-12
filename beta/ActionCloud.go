// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// CloudCommunicationsGetPresencesByUserIDRequestParameter undocumented
type CloudCommunicationsGetPresencesByUserIDRequestParameter struct {
	// IDs undocumented
	IDs []string `json:"ids,omitempty"`
}

// Calls returns request builder for Call collection
func (b *CloudCommunicationsRequestBuilder) Calls() *CloudCommunicationsCallsCollectionRequestBuilder {
	bb := &CloudCommunicationsCallsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calls"
	return bb
}

// CloudCommunicationsCallsCollectionRequestBuilder is request builder for Call collection
type CloudCommunicationsCallsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Call collection
func (b *CloudCommunicationsCallsCollectionRequestBuilder) Request() *CloudCommunicationsCallsCollectionRequest {
	return &CloudCommunicationsCallsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Call item
func (b *CloudCommunicationsCallsCollectionRequestBuilder) ID(id string) *CallRequestBuilder {
	bb := &CallRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CloudCommunicationsCallsCollectionRequest is request for Call collection
type CloudCommunicationsCallsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Call collection
func (r *CloudCommunicationsCallsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Call, error) {
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
	var values []Call
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
			value  []Call
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

// Get performs GET request for Call collection
func (r *CloudCommunicationsCallsCollectionRequest) Get(ctx context.Context) ([]Call, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Call collection
func (r *CloudCommunicationsCallsCollectionRequest) Add(ctx context.Context, reqObj *Call) (resObj *Call, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// OnlineMeetings returns request builder for OnlineMeeting collection
func (b *CloudCommunicationsRequestBuilder) OnlineMeetings() *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder {
	bb := &CloudCommunicationsOnlineMeetingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/onlineMeetings"
	return bb
}

// CloudCommunicationsOnlineMeetingsCollectionRequestBuilder is request builder for OnlineMeeting collection
type CloudCommunicationsOnlineMeetingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnlineMeeting collection
func (b *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder) Request() *CloudCommunicationsOnlineMeetingsCollectionRequest {
	return &CloudCommunicationsOnlineMeetingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnlineMeeting item
func (b *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder) ID(id string) *OnlineMeetingRequestBuilder {
	bb := &OnlineMeetingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CloudCommunicationsOnlineMeetingsCollectionRequest is request for OnlineMeeting collection
type CloudCommunicationsOnlineMeetingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnlineMeeting collection
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnlineMeeting, error) {
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
	var values []OnlineMeeting
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
			value  []OnlineMeeting
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

// Get performs GET request for OnlineMeeting collection
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) Get(ctx context.Context) ([]OnlineMeeting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnlineMeeting collection
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) Add(ctx context.Context, reqObj *OnlineMeeting) (resObj *OnlineMeeting, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
