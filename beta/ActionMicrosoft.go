// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ContainedApps returns request builder for MobileContainedApp collection
func (b *MicrosoftStoreForBusinessAppRequestBuilder) ContainedApps() *MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder {
	bb := &MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/containedApps"
	return bb
}

// MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder is request builder for MobileContainedApp collection
type MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileContainedApp collection
func (b *MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder) Request() *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest {
	return &MicrosoftStoreForBusinessAppContainedAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileContainedApp item
func (b *MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder) ID(id string) *MobileContainedAppRequestBuilder {
	bb := &MobileContainedAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MicrosoftStoreForBusinessAppContainedAppsCollectionRequest is request for MobileContainedApp collection
type MicrosoftStoreForBusinessAppContainedAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileContainedApp collection
func (r *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileContainedApp, error) {
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
	var values []MobileContainedApp
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
			value  []MobileContainedApp
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

// Get performs GET request for MobileContainedApp collection
func (r *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest) Get(ctx context.Context) ([]MobileContainedApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileContainedApp collection
func (r *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest) Add(ctx context.Context, reqObj *MobileContainedApp) (resObj *MobileContainedApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
