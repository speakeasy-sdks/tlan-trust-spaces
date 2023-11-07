// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"tlan-trust-spaces/v2/pkg/models/shared"
)

type ListConversationsBySpaceIDRequest struct {
	SpaceID string `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *ListConversationsBySpaceIDRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

type ListConversationsBySpaceIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// ok
	Inbox *shared.Inbox
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListConversationsBySpaceIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListConversationsBySpaceIDResponse) GetInbox() *shared.Inbox {
	if o == nil {
		return nil
	}
	return o.Inbox
}

func (o *ListConversationsBySpaceIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListConversationsBySpaceIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
