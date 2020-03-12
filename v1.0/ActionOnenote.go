// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OnenotePageOnenotePatchContentRequestParameter undocumented
type OnenotePageOnenotePatchContentRequestParameter struct {
	// Commands undocumented
	Commands []OnenotePatchContentCommand `json:"commands,omitempty"`
}

// OnenotePageCopyToSectionRequestParameter undocumented
type OnenotePageCopyToSectionRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
	// SiteCollectionID undocumented
	SiteCollectionID *string `json:"siteCollectionId,omitempty"`
	// SiteID undocumented
	SiteID *string `json:"siteId,omitempty"`
}

// OnenoteSectionCopyToNotebookRequestParameter undocumented
type OnenoteSectionCopyToNotebookRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
	// RenameAs undocumented
	RenameAs *string `json:"renameAs,omitempty"`
	// SiteCollectionID undocumented
	SiteCollectionID *string `json:"siteCollectionId,omitempty"`
	// SiteID undocumented
	SiteID *string `json:"siteId,omitempty"`
}

// OnenoteSectionCopyToSectionGroupRequestParameter undocumented
type OnenoteSectionCopyToSectionGroupRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
	// RenameAs undocumented
	RenameAs *string `json:"renameAs,omitempty"`
	// SiteCollectionID undocumented
	SiteCollectionID *string `json:"siteCollectionId,omitempty"`
	// SiteID undocumented
	SiteID *string `json:"siteId,omitempty"`
}

// Notebooks returns request builder for Notebook collection
func (b *OnenoteRequestBuilder) Notebooks() *OnenoteNotebooksCollectionRequestBuilder {
	bb := &OnenoteNotebooksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/notebooks"
	return bb
}

// OnenoteNotebooksCollectionRequestBuilder is request builder for Notebook collection
type OnenoteNotebooksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Notebook collection
func (b *OnenoteNotebooksCollectionRequestBuilder) Request() *OnenoteNotebooksCollectionRequest {
	return &OnenoteNotebooksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Notebook item
func (b *OnenoteNotebooksCollectionRequestBuilder) ID(id string) *NotebookRequestBuilder {
	bb := &NotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnenoteNotebooksCollectionRequest is request for Notebook collection
type OnenoteNotebooksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Notebook collection
func (r *OnenoteNotebooksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Notebook, error) {
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
	var values []Notebook
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
			value  []Notebook
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

// Get performs GET request for Notebook collection
func (r *OnenoteNotebooksCollectionRequest) Get(ctx context.Context) ([]Notebook, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Notebook collection
func (r *OnenoteNotebooksCollectionRequest) Add(ctx context.Context, reqObj *Notebook) (resObj *Notebook, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Operations returns request builder for OnenoteOperation collection
func (b *OnenoteRequestBuilder) Operations() *OnenoteOperationsCollectionRequestBuilder {
	bb := &OnenoteOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// OnenoteOperationsCollectionRequestBuilder is request builder for OnenoteOperation collection
type OnenoteOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenoteOperation collection
func (b *OnenoteOperationsCollectionRequestBuilder) Request() *OnenoteOperationsCollectionRequest {
	return &OnenoteOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenoteOperation item
func (b *OnenoteOperationsCollectionRequestBuilder) ID(id string) *OnenoteOperationRequestBuilder {
	bb := &OnenoteOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnenoteOperationsCollectionRequest is request for OnenoteOperation collection
type OnenoteOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnenoteOperation collection
func (r *OnenoteOperationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnenoteOperation, error) {
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
	var values []OnenoteOperation
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
			value  []OnenoteOperation
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

// Get performs GET request for OnenoteOperation collection
func (r *OnenoteOperationsCollectionRequest) Get(ctx context.Context) ([]OnenoteOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnenoteOperation collection
func (r *OnenoteOperationsCollectionRequest) Add(ctx context.Context, reqObj *OnenoteOperation) (resObj *OnenoteOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Pages returns request builder for OnenotePage collection
func (b *OnenoteRequestBuilder) Pages() *OnenotePagesCollectionRequestBuilder {
	bb := &OnenotePagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/pages"
	return bb
}

// OnenotePagesCollectionRequestBuilder is request builder for OnenotePage collection
type OnenotePagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenotePage collection
func (b *OnenotePagesCollectionRequestBuilder) Request() *OnenotePagesCollectionRequest {
	return &OnenotePagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenotePage item
func (b *OnenotePagesCollectionRequestBuilder) ID(id string) *OnenotePageRequestBuilder {
	bb := &OnenotePageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnenotePagesCollectionRequest is request for OnenotePage collection
type OnenotePagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnenotePage collection
func (r *OnenotePagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnenotePage, error) {
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
	var values []OnenotePage
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
			value  []OnenotePage
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

// Get performs GET request for OnenotePage collection
func (r *OnenotePagesCollectionRequest) Get(ctx context.Context) ([]OnenotePage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnenotePage collection
func (r *OnenotePagesCollectionRequest) Add(ctx context.Context, reqObj *OnenotePage) (resObj *OnenotePage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Resources returns request builder for OnenoteResource collection
func (b *OnenoteRequestBuilder) Resources() *OnenoteResourcesCollectionRequestBuilder {
	bb := &OnenoteResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resources"
	return bb
}

// OnenoteResourcesCollectionRequestBuilder is request builder for OnenoteResource collection
type OnenoteResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenoteResource collection
func (b *OnenoteResourcesCollectionRequestBuilder) Request() *OnenoteResourcesCollectionRequest {
	return &OnenoteResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenoteResource item
func (b *OnenoteResourcesCollectionRequestBuilder) ID(id string) *OnenoteResourceRequestBuilder {
	bb := &OnenoteResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnenoteResourcesCollectionRequest is request for OnenoteResource collection
type OnenoteResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnenoteResource collection
func (r *OnenoteResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnenoteResource, error) {
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
	var values []OnenoteResource
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
			value  []OnenoteResource
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

// Get performs GET request for OnenoteResource collection
func (r *OnenoteResourcesCollectionRequest) Get(ctx context.Context) ([]OnenoteResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnenoteResource collection
func (r *OnenoteResourcesCollectionRequest) Add(ctx context.Context, reqObj *OnenoteResource) (resObj *OnenoteResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SectionGroups returns request builder for SectionGroup collection
func (b *OnenoteRequestBuilder) SectionGroups() *OnenoteSectionGroupsCollectionRequestBuilder {
	bb := &OnenoteSectionGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sectionGroups"
	return bb
}

// OnenoteSectionGroupsCollectionRequestBuilder is request builder for SectionGroup collection
type OnenoteSectionGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SectionGroup collection
func (b *OnenoteSectionGroupsCollectionRequestBuilder) Request() *OnenoteSectionGroupsCollectionRequest {
	return &OnenoteSectionGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SectionGroup item
func (b *OnenoteSectionGroupsCollectionRequestBuilder) ID(id string) *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnenoteSectionGroupsCollectionRequest is request for SectionGroup collection
type OnenoteSectionGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SectionGroup collection
func (r *OnenoteSectionGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SectionGroup, error) {
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
	var values []SectionGroup
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
			value  []SectionGroup
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

// Get performs GET request for SectionGroup collection
func (r *OnenoteSectionGroupsCollectionRequest) Get(ctx context.Context) ([]SectionGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SectionGroup collection
func (r *OnenoteSectionGroupsCollectionRequest) Add(ctx context.Context, reqObj *SectionGroup) (resObj *SectionGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Sections returns request builder for OnenoteSection collection
func (b *OnenoteRequestBuilder) Sections() *OnenoteSectionsCollectionRequestBuilder {
	bb := &OnenoteSectionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sections"
	return bb
}

// OnenoteSectionsCollectionRequestBuilder is request builder for OnenoteSection collection
type OnenoteSectionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenoteSection collection
func (b *OnenoteSectionsCollectionRequestBuilder) Request() *OnenoteSectionsCollectionRequest {
	return &OnenoteSectionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenoteSection item
func (b *OnenoteSectionsCollectionRequestBuilder) ID(id string) *OnenoteSectionRequestBuilder {
	bb := &OnenoteSectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnenoteSectionsCollectionRequest is request for OnenoteSection collection
type OnenoteSectionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnenoteSection collection
func (r *OnenoteSectionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnenoteSection, error) {
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
	var values []OnenoteSection
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
			value  []OnenoteSection
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

// Get performs GET request for OnenoteSection collection
func (r *OnenoteSectionsCollectionRequest) Get(ctx context.Context) ([]OnenoteSection, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnenoteSection collection
func (r *OnenoteSectionsCollectionRequest) Add(ctx context.Context, reqObj *OnenoteSection) (resObj *OnenoteSection, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ParentNotebook is navigation property
func (b *OnenotePageRequestBuilder) ParentNotebook() *NotebookRequestBuilder {
	bb := &NotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentNotebook"
	return bb
}

// ParentSection is navigation property
func (b *OnenotePageRequestBuilder) ParentSection() *OnenoteSectionRequestBuilder {
	bb := &OnenoteSectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentSection"
	return bb
}

// Pages returns request builder for OnenotePage collection
func (b *OnenoteSectionRequestBuilder) Pages() *OnenoteSectionPagesCollectionRequestBuilder {
	bb := &OnenoteSectionPagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/pages"
	return bb
}

// OnenoteSectionPagesCollectionRequestBuilder is request builder for OnenotePage collection
type OnenoteSectionPagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenotePage collection
func (b *OnenoteSectionPagesCollectionRequestBuilder) Request() *OnenoteSectionPagesCollectionRequest {
	return &OnenoteSectionPagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenotePage item
func (b *OnenoteSectionPagesCollectionRequestBuilder) ID(id string) *OnenotePageRequestBuilder {
	bb := &OnenotePageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnenoteSectionPagesCollectionRequest is request for OnenotePage collection
type OnenoteSectionPagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnenotePage collection
func (r *OnenoteSectionPagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnenotePage, error) {
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
	var values []OnenotePage
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
			value  []OnenotePage
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

// Get performs GET request for OnenotePage collection
func (r *OnenoteSectionPagesCollectionRequest) Get(ctx context.Context) ([]OnenotePage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnenotePage collection
func (r *OnenoteSectionPagesCollectionRequest) Add(ctx context.Context, reqObj *OnenotePage) (resObj *OnenotePage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ParentNotebook is navigation property
func (b *OnenoteSectionRequestBuilder) ParentNotebook() *NotebookRequestBuilder {
	bb := &NotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentNotebook"
	return bb
}

// ParentSectionGroup is navigation property
func (b *OnenoteSectionRequestBuilder) ParentSectionGroup() *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentSectionGroup"
	return bb
}
