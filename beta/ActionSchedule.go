// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ScheduleShareRequestParameter undocumented
type ScheduleShareRequestParameter struct {
	// NotifyTeam undocumented
	NotifyTeam *bool `json:"notifyTeam,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
}

// ScheduleChangeRequestObjectApproveRequestParameter undocumented
type ScheduleChangeRequestObjectApproveRequestParameter struct {
	// Message undocumented
	Message *string `json:"message,omitempty"`
}

// ScheduleChangeRequestObjectDeclineRequestParameter undocumented
type ScheduleChangeRequestObjectDeclineRequestParameter struct {
	// Message undocumented
	Message *string `json:"message,omitempty"`
}

// OpenShiftChangeRequests returns request builder for OpenShiftChangeRequestObject collection
func (b *ScheduleRequestBuilder) OpenShiftChangeRequests() *ScheduleOpenShiftChangeRequestsCollectionRequestBuilder {
	bb := &ScheduleOpenShiftChangeRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/openShiftChangeRequests"
	return bb
}

// ScheduleOpenShiftChangeRequestsCollectionRequestBuilder is request builder for OpenShiftChangeRequestObject collection
type ScheduleOpenShiftChangeRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OpenShiftChangeRequestObject collection
func (b *ScheduleOpenShiftChangeRequestsCollectionRequestBuilder) Request() *ScheduleOpenShiftChangeRequestsCollectionRequest {
	return &ScheduleOpenShiftChangeRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OpenShiftChangeRequestObject item
func (b *ScheduleOpenShiftChangeRequestsCollectionRequestBuilder) ID(id string) *OpenShiftChangeRequestObjectRequestBuilder {
	bb := &OpenShiftChangeRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ScheduleOpenShiftChangeRequestsCollectionRequest is request for OpenShiftChangeRequestObject collection
type ScheduleOpenShiftChangeRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OpenShiftChangeRequestObject collection
func (r *ScheduleOpenShiftChangeRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OpenShiftChangeRequestObject, error) {
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
	var values []OpenShiftChangeRequestObject
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
			value  []OpenShiftChangeRequestObject
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

// Get performs GET request for OpenShiftChangeRequestObject collection
func (r *ScheduleOpenShiftChangeRequestsCollectionRequest) Get(ctx context.Context) ([]OpenShiftChangeRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OpenShiftChangeRequestObject collection
func (r *ScheduleOpenShiftChangeRequestsCollectionRequest) Add(ctx context.Context, reqObj *OpenShiftChangeRequestObject) (resObj *OpenShiftChangeRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// OpenShifts returns request builder for OpenShift collection
func (b *ScheduleRequestBuilder) OpenShifts() *ScheduleOpenShiftsCollectionRequestBuilder {
	bb := &ScheduleOpenShiftsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/openShifts"
	return bb
}

// ScheduleOpenShiftsCollectionRequestBuilder is request builder for OpenShift collection
type ScheduleOpenShiftsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OpenShift collection
func (b *ScheduleOpenShiftsCollectionRequestBuilder) Request() *ScheduleOpenShiftsCollectionRequest {
	return &ScheduleOpenShiftsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OpenShift item
func (b *ScheduleOpenShiftsCollectionRequestBuilder) ID(id string) *OpenShiftRequestBuilder {
	bb := &OpenShiftRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ScheduleOpenShiftsCollectionRequest is request for OpenShift collection
type ScheduleOpenShiftsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OpenShift collection
func (r *ScheduleOpenShiftsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OpenShift, error) {
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
	var values []OpenShift
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
			value  []OpenShift
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

// Get performs GET request for OpenShift collection
func (r *ScheduleOpenShiftsCollectionRequest) Get(ctx context.Context) ([]OpenShift, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OpenShift collection
func (r *ScheduleOpenShiftsCollectionRequest) Add(ctx context.Context, reqObj *OpenShift) (resObj *OpenShift, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SchedulingGroups returns request builder for SchedulingGroup collection
func (b *ScheduleRequestBuilder) SchedulingGroups() *ScheduleSchedulingGroupsCollectionRequestBuilder {
	bb := &ScheduleSchedulingGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schedulingGroups"
	return bb
}

// ScheduleSchedulingGroupsCollectionRequestBuilder is request builder for SchedulingGroup collection
type ScheduleSchedulingGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SchedulingGroup collection
func (b *ScheduleSchedulingGroupsCollectionRequestBuilder) Request() *ScheduleSchedulingGroupsCollectionRequest {
	return &ScheduleSchedulingGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SchedulingGroup item
func (b *ScheduleSchedulingGroupsCollectionRequestBuilder) ID(id string) *SchedulingGroupRequestBuilder {
	bb := &SchedulingGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ScheduleSchedulingGroupsCollectionRequest is request for SchedulingGroup collection
type ScheduleSchedulingGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SchedulingGroup collection
func (r *ScheduleSchedulingGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SchedulingGroup, error) {
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
	var values []SchedulingGroup
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
			value  []SchedulingGroup
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

// Get performs GET request for SchedulingGroup collection
func (r *ScheduleSchedulingGroupsCollectionRequest) Get(ctx context.Context) ([]SchedulingGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SchedulingGroup collection
func (r *ScheduleSchedulingGroupsCollectionRequest) Add(ctx context.Context, reqObj *SchedulingGroup) (resObj *SchedulingGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Shifts returns request builder for Shift collection
func (b *ScheduleRequestBuilder) Shifts() *ScheduleShiftsCollectionRequestBuilder {
	bb := &ScheduleShiftsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/shifts"
	return bb
}

// ScheduleShiftsCollectionRequestBuilder is request builder for Shift collection
type ScheduleShiftsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Shift collection
func (b *ScheduleShiftsCollectionRequestBuilder) Request() *ScheduleShiftsCollectionRequest {
	return &ScheduleShiftsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Shift item
func (b *ScheduleShiftsCollectionRequestBuilder) ID(id string) *ShiftRequestBuilder {
	bb := &ShiftRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ScheduleShiftsCollectionRequest is request for Shift collection
type ScheduleShiftsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Shift collection
func (r *ScheduleShiftsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Shift, error) {
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
	var values []Shift
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
			value  []Shift
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

// Get performs GET request for Shift collection
func (r *ScheduleShiftsCollectionRequest) Get(ctx context.Context) ([]Shift, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Shift collection
func (r *ScheduleShiftsCollectionRequest) Add(ctx context.Context, reqObj *Shift) (resObj *Shift, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SwapShiftsChangeRequests returns request builder for SwapShiftsChangeRequestObject collection
func (b *ScheduleRequestBuilder) SwapShiftsChangeRequests() *ScheduleSwapShiftsChangeRequestsCollectionRequestBuilder {
	bb := &ScheduleSwapShiftsChangeRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/swapShiftsChangeRequests"
	return bb
}

// ScheduleSwapShiftsChangeRequestsCollectionRequestBuilder is request builder for SwapShiftsChangeRequestObject collection
type ScheduleSwapShiftsChangeRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SwapShiftsChangeRequestObject collection
func (b *ScheduleSwapShiftsChangeRequestsCollectionRequestBuilder) Request() *ScheduleSwapShiftsChangeRequestsCollectionRequest {
	return &ScheduleSwapShiftsChangeRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SwapShiftsChangeRequestObject item
func (b *ScheduleSwapShiftsChangeRequestsCollectionRequestBuilder) ID(id string) *SwapShiftsChangeRequestObjectRequestBuilder {
	bb := &SwapShiftsChangeRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ScheduleSwapShiftsChangeRequestsCollectionRequest is request for SwapShiftsChangeRequestObject collection
type ScheduleSwapShiftsChangeRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SwapShiftsChangeRequestObject collection
func (r *ScheduleSwapShiftsChangeRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SwapShiftsChangeRequestObject, error) {
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
	var values []SwapShiftsChangeRequestObject
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
			value  []SwapShiftsChangeRequestObject
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

// Get performs GET request for SwapShiftsChangeRequestObject collection
func (r *ScheduleSwapShiftsChangeRequestsCollectionRequest) Get(ctx context.Context) ([]SwapShiftsChangeRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SwapShiftsChangeRequestObject collection
func (r *ScheduleSwapShiftsChangeRequestsCollectionRequest) Add(ctx context.Context, reqObj *SwapShiftsChangeRequestObject) (resObj *SwapShiftsChangeRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TimeOffReasons returns request builder for TimeOffReason collection
func (b *ScheduleRequestBuilder) TimeOffReasons() *ScheduleTimeOffReasonsCollectionRequestBuilder {
	bb := &ScheduleTimeOffReasonsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/timeOffReasons"
	return bb
}

// ScheduleTimeOffReasonsCollectionRequestBuilder is request builder for TimeOffReason collection
type ScheduleTimeOffReasonsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TimeOffReason collection
func (b *ScheduleTimeOffReasonsCollectionRequestBuilder) Request() *ScheduleTimeOffReasonsCollectionRequest {
	return &ScheduleTimeOffReasonsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TimeOffReason item
func (b *ScheduleTimeOffReasonsCollectionRequestBuilder) ID(id string) *TimeOffReasonRequestBuilder {
	bb := &TimeOffReasonRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ScheduleTimeOffReasonsCollectionRequest is request for TimeOffReason collection
type ScheduleTimeOffReasonsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TimeOffReason collection
func (r *ScheduleTimeOffReasonsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]TimeOffReason, error) {
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
	var values []TimeOffReason
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
			value  []TimeOffReason
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

// Get performs GET request for TimeOffReason collection
func (r *ScheduleTimeOffReasonsCollectionRequest) Get(ctx context.Context) ([]TimeOffReason, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for TimeOffReason collection
func (r *ScheduleTimeOffReasonsCollectionRequest) Add(ctx context.Context, reqObj *TimeOffReason) (resObj *TimeOffReason, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TimeOffRequests returns request builder for TimeOffRequestObject collection
func (b *ScheduleRequestBuilder) TimeOffRequests() *ScheduleTimeOffRequestsCollectionRequestBuilder {
	bb := &ScheduleTimeOffRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/timeOffRequests"
	return bb
}

// ScheduleTimeOffRequestsCollectionRequestBuilder is request builder for TimeOffRequestObject collection
type ScheduleTimeOffRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TimeOffRequestObject collection
func (b *ScheduleTimeOffRequestsCollectionRequestBuilder) Request() *ScheduleTimeOffRequestsCollectionRequest {
	return &ScheduleTimeOffRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TimeOffRequestObject item
func (b *ScheduleTimeOffRequestsCollectionRequestBuilder) ID(id string) *TimeOffRequestObjectRequestBuilder {
	bb := &TimeOffRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ScheduleTimeOffRequestsCollectionRequest is request for TimeOffRequestObject collection
type ScheduleTimeOffRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TimeOffRequestObject collection
func (r *ScheduleTimeOffRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]TimeOffRequestObject, error) {
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
	var values []TimeOffRequestObject
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
			value  []TimeOffRequestObject
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

// Get performs GET request for TimeOffRequestObject collection
func (r *ScheduleTimeOffRequestsCollectionRequest) Get(ctx context.Context) ([]TimeOffRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for TimeOffRequestObject collection
func (r *ScheduleTimeOffRequestsCollectionRequest) Add(ctx context.Context, reqObj *TimeOffRequestObject) (resObj *TimeOffRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TimesOff returns request builder for TimeOff collection
func (b *ScheduleRequestBuilder) TimesOff() *ScheduleTimesOffCollectionRequestBuilder {
	bb := &ScheduleTimesOffCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/timesOff"
	return bb
}

// ScheduleTimesOffCollectionRequestBuilder is request builder for TimeOff collection
type ScheduleTimesOffCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TimeOff collection
func (b *ScheduleTimesOffCollectionRequestBuilder) Request() *ScheduleTimesOffCollectionRequest {
	return &ScheduleTimesOffCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TimeOff item
func (b *ScheduleTimesOffCollectionRequestBuilder) ID(id string) *TimeOffRequestBuilder {
	bb := &TimeOffRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ScheduleTimesOffCollectionRequest is request for TimeOff collection
type ScheduleTimesOffCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TimeOff collection
func (r *ScheduleTimesOffCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]TimeOff, error) {
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
	var values []TimeOff
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
			value  []TimeOff
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

// Get performs GET request for TimeOff collection
func (r *ScheduleTimesOffCollectionRequest) Get(ctx context.Context) ([]TimeOff, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for TimeOff collection
func (r *ScheduleTimesOffCollectionRequest) Add(ctx context.Context, reqObj *TimeOff) (resObj *TimeOff, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
