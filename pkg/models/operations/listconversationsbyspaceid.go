// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"tlan-trust-spaces/pkg/models/shared"
)

type ListConversationsBySpaceIDRequest struct {
	SpaceID string `pathParam:"style=simple,explode=false,name=spaceId"`
}

type ListConversationsBySpaceIDResponse struct {
	ContentType string
	// ok
	Inbox       *shared.Inbox
	StatusCode  int
	RawResponse *http.Response
}