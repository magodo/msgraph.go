// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SearchRequestBuilder is request builder for Search
type SearchRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchRequest
func (b *SearchRequestBuilder) Request() *SearchRequest {
	return &SearchRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchRequest is request for Search
type SearchRequest struct{ BaseRequest }

// Get performs GET request for Search
func (r *SearchRequest) Get(ctx context.Context) (resObj *Search, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Search
func (r *SearchRequest) Update(ctx context.Context, reqObj *Search) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Search
func (r *SearchRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type SearchQueryRequestBuilder struct{ BaseRequestBuilder }

// Query action undocumented
func (b *SearchRequestBuilder) Query(reqObj *SearchQueryRequestParameter) *SearchQueryRequestBuilder {
	bb := &SearchQueryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/query"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SearchQueryRequest struct{ BaseRequest }

//
func (b *SearchQueryRequestBuilder) Request() *SearchQueryRequest {
	return &SearchQueryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SearchQueryRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SearchResponse, error) {
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
	var values []SearchResponse
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
			value  []SearchResponse
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

//
func (r *SearchQueryRequest) PostN(ctx context.Context, n int) ([]SearchResponse, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *SearchQueryRequest) Post(ctx context.Context) ([]SearchResponse, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}
