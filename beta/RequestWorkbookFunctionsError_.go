// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsError_TypeRequestBuilder struct{ BaseRequestBuilder }

// Error_Type action undocumented
func (b *WorkbookFunctionsRequestBuilder) Error_Type(reqObj *WorkbookFunctionsError_TypeRequestParameter) *WorkbookFunctionsError_TypeRequestBuilder {
	bb := &WorkbookFunctionsError_TypeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/error_Type"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsError_TypeRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsError_TypeRequestBuilder) Request() *WorkbookFunctionsError_TypeRequest {
	return &WorkbookFunctionsError_TypeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsError_TypeRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
