// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"tlan-trust-spaces/pkg/models/shared"
)

type GetConversationByIDRequest struct {
	ConversationID string `pathParam:"style=simple,explode=false,name=conversationId"`
}

func (o *GetConversationByIDRequest) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

type GetConversationByIDResponse struct {
	ContentType string
	// Conversation details
	Conversation *shared.Conversation
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetConversationByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConversationByIDResponse) GetConversation() *shared.Conversation {
	if o == nil {
		return nil
	}
	return o.Conversation
}

func (o *GetConversationByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConversationByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
