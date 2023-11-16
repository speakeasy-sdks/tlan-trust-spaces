// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"tlan-trust-spaces/v2/pkg/models/shared"
)

type StartConversationRequest struct {
	ConversationStart *shared.ConversationStart `request:"mediaType=application/json"`
	SpaceID           string                    `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *StartConversationRequest) GetConversationStart() *shared.ConversationStart {
	if o == nil {
		return nil
	}
	return o.ConversationStart
}

func (o *StartConversationRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

type StartConversationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// ok
	ConversationStartResponse *shared.ConversationStartResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *StartConversationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StartConversationResponse) GetConversationStartResponse() *shared.ConversationStartResponse {
	if o == nil {
		return nil
	}
	return o.ConversationStartResponse
}

func (o *StartConversationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StartConversationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
