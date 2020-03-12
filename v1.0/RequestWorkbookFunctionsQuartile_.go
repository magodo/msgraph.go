// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsQuartile_ExcRequestBuilder struct{ BaseRequestBuilder }

// Quartile_Exc action undocumented
func (b *WorkbookFunctionsRequestBuilder) Quartile_Exc(reqObj *WorkbookFunctionsQuartile_ExcRequestParameter) *WorkbookFunctionsQuartile_ExcRequestBuilder {
	bb := &WorkbookFunctionsQuartile_ExcRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/quartile_Exc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsQuartile_ExcRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsQuartile_ExcRequestBuilder) Request() *WorkbookFunctionsQuartile_ExcRequest {
	return &WorkbookFunctionsQuartile_ExcRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsQuartile_ExcRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsQuartile_IncRequestBuilder struct{ BaseRequestBuilder }

// Quartile_Inc action undocumented
func (b *WorkbookFunctionsRequestBuilder) Quartile_Inc(reqObj *WorkbookFunctionsQuartile_IncRequestParameter) *WorkbookFunctionsQuartile_IncRequestBuilder {
	bb := &WorkbookFunctionsQuartile_IncRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/quartile_Inc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsQuartile_IncRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsQuartile_IncRequestBuilder) Request() *WorkbookFunctionsQuartile_IncRequest {
	return &WorkbookFunctionsQuartile_IncRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsQuartile_IncRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
