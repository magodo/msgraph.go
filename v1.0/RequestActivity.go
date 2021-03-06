// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ActivityHistoryItemRequestBuilder is request builder for ActivityHistoryItem
type ActivityHistoryItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityHistoryItemRequest
func (b *ActivityHistoryItemRequestBuilder) Request() *ActivityHistoryItemRequest {
	return &ActivityHistoryItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActivityHistoryItemRequest is request for ActivityHistoryItem
type ActivityHistoryItemRequest struct{ BaseRequest }

// Get performs GET request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Get(ctx context.Context) (resObj *ActivityHistoryItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Update(ctx context.Context, reqObj *ActivityHistoryItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
