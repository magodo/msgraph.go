// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Resource is navigation property
func (b *UsedInsightRequestBuilder) Resource() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}
