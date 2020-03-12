// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SiteCollectionAddRequestParameter undocumented
type SiteCollectionAddRequestParameter struct {
	// Value undocumented
	Value []Site `json:"value,omitempty"`
}

// SiteCollectionRemoveRequestParameter undocumented
type SiteCollectionRemoveRequestParameter struct {
	// Value undocumented
	Value []Site `json:"value,omitempty"`
}

// SitePagePublishRequestParameter undocumented
type SitePagePublishRequestParameter struct {
}

// Analytics is navigation property
func (b *SiteRequestBuilder) Analytics() *ItemAnalyticsRequestBuilder {
	bb := &ItemAnalyticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/analytics"
	return bb
}

// Columns returns request builder for ColumnDefinition collection
func (b *SiteRequestBuilder) Columns() *SiteColumnsCollectionRequestBuilder {
	bb := &SiteColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columns"
	return bb
}

// SiteColumnsCollectionRequestBuilder is request builder for ColumnDefinition collection
type SiteColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *SiteColumnsCollectionRequestBuilder) Request() *SiteColumnsCollectionRequest {
	return &SiteColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *SiteColumnsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteColumnsCollectionRequest is request for ColumnDefinition collection
type SiteColumnsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnDefinition collection
func (r *SiteColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ColumnDefinition, error) {
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
	var values []ColumnDefinition
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
			value  []ColumnDefinition
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

// Get performs GET request for ColumnDefinition collection
func (r *SiteColumnsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ColumnDefinition collection
func (r *SiteColumnsCollectionRequest) Add(ctx context.Context, reqObj *ColumnDefinition) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ContentTypes returns request builder for ContentType collection
func (b *SiteRequestBuilder) ContentTypes() *SiteContentTypesCollectionRequestBuilder {
	bb := &SiteContentTypesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/contentTypes"
	return bb
}

// SiteContentTypesCollectionRequestBuilder is request builder for ContentType collection
type SiteContentTypesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ContentType collection
func (b *SiteContentTypesCollectionRequestBuilder) Request() *SiteContentTypesCollectionRequest {
	return &SiteContentTypesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ContentType item
func (b *SiteContentTypesCollectionRequestBuilder) ID(id string) *ContentTypeRequestBuilder {
	bb := &ContentTypeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteContentTypesCollectionRequest is request for ContentType collection
type SiteContentTypesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ContentType collection
func (r *SiteContentTypesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ContentType, error) {
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
	var values []ContentType
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
			value  []ContentType
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

// Get performs GET request for ContentType collection
func (r *SiteContentTypesCollectionRequest) Get(ctx context.Context) ([]ContentType, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ContentType collection
func (r *SiteContentTypesCollectionRequest) Add(ctx context.Context, reqObj *ContentType) (resObj *ContentType, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Drive is navigation property
func (b *SiteRequestBuilder) Drive() *DriveRequestBuilder {
	bb := &DriveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/drive"
	return bb
}

// Drives returns request builder for Drive collection
func (b *SiteRequestBuilder) Drives() *SiteDrivesCollectionRequestBuilder {
	bb := &SiteDrivesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/drives"
	return bb
}

// SiteDrivesCollectionRequestBuilder is request builder for Drive collection
type SiteDrivesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Drive collection
func (b *SiteDrivesCollectionRequestBuilder) Request() *SiteDrivesCollectionRequest {
	return &SiteDrivesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Drive item
func (b *SiteDrivesCollectionRequestBuilder) ID(id string) *DriveRequestBuilder {
	bb := &DriveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteDrivesCollectionRequest is request for Drive collection
type SiteDrivesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Drive collection
func (r *SiteDrivesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Drive, error) {
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
	var values []Drive
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
			value  []Drive
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

// Get performs GET request for Drive collection
func (r *SiteDrivesCollectionRequest) Get(ctx context.Context) ([]Drive, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Drive collection
func (r *SiteDrivesCollectionRequest) Add(ctx context.Context, reqObj *Drive) (resObj *Drive, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Items returns request builder for BaseItem collection
func (b *SiteRequestBuilder) Items() *SiteItemsCollectionRequestBuilder {
	bb := &SiteItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// SiteItemsCollectionRequestBuilder is request builder for BaseItem collection
type SiteItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BaseItem collection
func (b *SiteItemsCollectionRequestBuilder) Request() *SiteItemsCollectionRequest {
	return &SiteItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BaseItem item
func (b *SiteItemsCollectionRequestBuilder) ID(id string) *BaseItemRequestBuilder {
	bb := &BaseItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteItemsCollectionRequest is request for BaseItem collection
type SiteItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BaseItem collection
func (r *SiteItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]BaseItem, error) {
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
	var values []BaseItem
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
			value  []BaseItem
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

// Get performs GET request for BaseItem collection
func (r *SiteItemsCollectionRequest) Get(ctx context.Context) ([]BaseItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for BaseItem collection
func (r *SiteItemsCollectionRequest) Add(ctx context.Context, reqObj *BaseItem) (resObj *BaseItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Lists returns request builder for List collection
func (b *SiteRequestBuilder) Lists() *SiteListsCollectionRequestBuilder {
	bb := &SiteListsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/lists"
	return bb
}

// SiteListsCollectionRequestBuilder is request builder for List collection
type SiteListsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for List collection
func (b *SiteListsCollectionRequestBuilder) Request() *SiteListsCollectionRequest {
	return &SiteListsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for List item
func (b *SiteListsCollectionRequestBuilder) ID(id string) *ListRequestBuilder {
	bb := &ListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteListsCollectionRequest is request for List collection
type SiteListsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for List collection
func (r *SiteListsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]List, error) {
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
	var values []List
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
			value  []List
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

// Get performs GET request for List collection
func (r *SiteListsCollectionRequest) Get(ctx context.Context) ([]List, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for List collection
func (r *SiteListsCollectionRequest) Add(ctx context.Context, reqObj *List) (resObj *List, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Onenote is navigation property
func (b *SiteRequestBuilder) Onenote() *OnenoteRequestBuilder {
	bb := &OnenoteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/onenote"
	return bb
}

// Pages returns request builder for SitePage collection
func (b *SiteRequestBuilder) Pages() *SitePagesCollectionRequestBuilder {
	bb := &SitePagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/pages"
	return bb
}

// SitePagesCollectionRequestBuilder is request builder for SitePage collection
type SitePagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SitePage collection
func (b *SitePagesCollectionRequestBuilder) Request() *SitePagesCollectionRequest {
	return &SitePagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SitePage item
func (b *SitePagesCollectionRequestBuilder) ID(id string) *SitePageRequestBuilder {
	bb := &SitePageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SitePagesCollectionRequest is request for SitePage collection
type SitePagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SitePage collection
func (r *SitePagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SitePage, error) {
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
	var values []SitePage
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
			value  []SitePage
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

// Get performs GET request for SitePage collection
func (r *SitePagesCollectionRequest) Get(ctx context.Context) ([]SitePage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SitePage collection
func (r *SitePagesCollectionRequest) Add(ctx context.Context, reqObj *SitePage) (resObj *SitePage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Sites returns request builder for Site collection
func (b *SiteRequestBuilder) Sites() *SiteSitesCollectionRequestBuilder {
	bb := &SiteSitesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sites"
	return bb
}

// SiteSitesCollectionRequestBuilder is request builder for Site collection
type SiteSitesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Site collection
func (b *SiteSitesCollectionRequestBuilder) Request() *SiteSitesCollectionRequest {
	return &SiteSitesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Site item
func (b *SiteSitesCollectionRequestBuilder) ID(id string) *SiteRequestBuilder {
	bb := &SiteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteSitesCollectionRequest is request for Site collection
type SiteSitesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Site collection
func (r *SiteSitesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Site, error) {
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
	var values []Site
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
			value  []Site
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

// Get performs GET request for Site collection
func (r *SiteSitesCollectionRequest) Get(ctx context.Context) ([]Site, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Site collection
func (r *SiteSitesCollectionRequest) Add(ctx context.Context, reqObj *Site) (resObj *Site, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
