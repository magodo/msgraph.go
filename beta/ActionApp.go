// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AppLogCollectionRequestObjectCreateDownloadURLRequestParameter undocumented
type AppLogCollectionRequestObjectCreateDownloadURLRequestParameter struct {
}

// TeamsApps returns request builder for TeamsApp collection
func (b *AppCatalogsRequestBuilder) TeamsApps() *AppCatalogsTeamsAppsCollectionRequestBuilder {
	bb := &AppCatalogsTeamsAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamsApps"
	return bb
}

// AppCatalogsTeamsAppsCollectionRequestBuilder is request builder for TeamsApp collection
type AppCatalogsTeamsAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsApp collection
func (b *AppCatalogsTeamsAppsCollectionRequestBuilder) Request() *AppCatalogsTeamsAppsCollectionRequest {
	return &AppCatalogsTeamsAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsApp item
func (b *AppCatalogsTeamsAppsCollectionRequestBuilder) ID(id string) *TeamsAppRequestBuilder {
	bb := &TeamsAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppCatalogsTeamsAppsCollectionRequest is request for TeamsApp collection
type AppCatalogsTeamsAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsApp collection
func (r *AppCatalogsTeamsAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsApp, error) {
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
	var values []TeamsApp
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
			value  []TeamsApp
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

// GetN performs GET request for TeamsApp collection, max N pages
func (r *AppCatalogsTeamsAppsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsApp collection
func (r *AppCatalogsTeamsAppsCollectionRequest) Get(ctx context.Context) ([]TeamsApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsApp collection
func (r *AppCatalogsTeamsAppsCollectionRequest) Add(ctx context.Context, reqObj *TeamsApp) (resObj *TeamsApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ManagedDevices returns request builder for AppVulnerabilityManagedDevice collection
func (b *AppVulnerabilityTaskRequestBuilder) ManagedDevices() *AppVulnerabilityTaskManagedDevicesCollectionRequestBuilder {
	bb := &AppVulnerabilityTaskManagedDevicesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDevices"
	return bb
}

// AppVulnerabilityTaskManagedDevicesCollectionRequestBuilder is request builder for AppVulnerabilityManagedDevice collection
type AppVulnerabilityTaskManagedDevicesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppVulnerabilityManagedDevice collection
func (b *AppVulnerabilityTaskManagedDevicesCollectionRequestBuilder) Request() *AppVulnerabilityTaskManagedDevicesCollectionRequest {
	return &AppVulnerabilityTaskManagedDevicesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppVulnerabilityManagedDevice item
func (b *AppVulnerabilityTaskManagedDevicesCollectionRequestBuilder) ID(id string) *AppVulnerabilityManagedDeviceRequestBuilder {
	bb := &AppVulnerabilityManagedDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppVulnerabilityTaskManagedDevicesCollectionRequest is request for AppVulnerabilityManagedDevice collection
type AppVulnerabilityTaskManagedDevicesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppVulnerabilityManagedDevice collection
func (r *AppVulnerabilityTaskManagedDevicesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AppVulnerabilityManagedDevice, error) {
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
	var values []AppVulnerabilityManagedDevice
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
			value  []AppVulnerabilityManagedDevice
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

// GetN performs GET request for AppVulnerabilityManagedDevice collection, max N pages
func (r *AppVulnerabilityTaskManagedDevicesCollectionRequest) GetN(ctx context.Context, n int) ([]AppVulnerabilityManagedDevice, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AppVulnerabilityManagedDevice collection
func (r *AppVulnerabilityTaskManagedDevicesCollectionRequest) Get(ctx context.Context) ([]AppVulnerabilityManagedDevice, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AppVulnerabilityManagedDevice collection
func (r *AppVulnerabilityTaskManagedDevicesCollectionRequest) Add(ctx context.Context, reqObj *AppVulnerabilityManagedDevice) (resObj *AppVulnerabilityManagedDevice, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MobileApps returns request builder for AppVulnerabilityMobileApp collection
func (b *AppVulnerabilityTaskRequestBuilder) MobileApps() *AppVulnerabilityTaskMobileAppsCollectionRequestBuilder {
	bb := &AppVulnerabilityTaskMobileAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/mobileApps"
	return bb
}

// AppVulnerabilityTaskMobileAppsCollectionRequestBuilder is request builder for AppVulnerabilityMobileApp collection
type AppVulnerabilityTaskMobileAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppVulnerabilityMobileApp collection
func (b *AppVulnerabilityTaskMobileAppsCollectionRequestBuilder) Request() *AppVulnerabilityTaskMobileAppsCollectionRequest {
	return &AppVulnerabilityTaskMobileAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppVulnerabilityMobileApp item
func (b *AppVulnerabilityTaskMobileAppsCollectionRequestBuilder) ID(id string) *AppVulnerabilityMobileAppRequestBuilder {
	bb := &AppVulnerabilityMobileAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppVulnerabilityTaskMobileAppsCollectionRequest is request for AppVulnerabilityMobileApp collection
type AppVulnerabilityTaskMobileAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppVulnerabilityMobileApp collection
func (r *AppVulnerabilityTaskMobileAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AppVulnerabilityMobileApp, error) {
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
	var values []AppVulnerabilityMobileApp
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
			value  []AppVulnerabilityMobileApp
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

// GetN performs GET request for AppVulnerabilityMobileApp collection, max N pages
func (r *AppVulnerabilityTaskMobileAppsCollectionRequest) GetN(ctx context.Context, n int) ([]AppVulnerabilityMobileApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AppVulnerabilityMobileApp collection
func (r *AppVulnerabilityTaskMobileAppsCollectionRequest) Get(ctx context.Context) ([]AppVulnerabilityMobileApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AppVulnerabilityMobileApp collection
func (r *AppVulnerabilityTaskMobileAppsCollectionRequest) Add(ctx context.Context, reqObj *AppVulnerabilityMobileApp) (resObj *AppVulnerabilityMobileApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
